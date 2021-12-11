package service

import (
	"chat-room/internal/dao/pool"
	"chat-room/pkg/common/constant"
	"chat-room/pkg/common/response"
	"chat-room/pkg/errors"
	"chat-room/pkg/global/log"
	"chat-room/pkg/protocol"

	"chat-room/internal/model"
	"chat-room/pkg/common/request"

	"gorm.io/gorm"
)

const NULL_ID int32 = 0

type messageService struct {
}

var MessageService = new(messageService)

func (m *messageService) GetMessages(message request.MessageRequest) ([]response.MessageResponse, error) {
	db := pool.GetDB()

	migrate := &model.Message{}
	pool.GetDB().AutoMigrate(&migrate)

	if message.MessageType == constant.MESSAGE_TYPE_USER {
		var queryUser *model.User
		db.First(&queryUser, "uuid = ?", message.Uuid)

		if NULL_ID == queryUser.Id {
			return nil, errors.New("用户不存在")
		}

		var friend *model.User
		db.First(&friend, "username = ?", message.FriendUsername)
		if NULL_ID == friend.Id {
			return nil, errors.New("用户不存在")
		}

		var messages []response.MessageResponse

		db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar, to_user.username AS to_username  FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id LEFT JOIN users AS to_user ON m.to_user_id = to_user.id WHERE from_user_id IN (?, ?) AND to_user_id IN (?, ?)",
			queryUser.Id, friend.Id, queryUser.Id, friend.Id).Scan(&messages)

		return messages, nil
	}

	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		messages, err := fetchGroupMessage(db, message.Uuid)
		if err != nil {
			return nil, err
		}

		return messages, nil
	}

	return nil, errors.New("不支持查询类型")
}

func fetchGroupMessage(db *gorm.DB, toUuid string) ([]response.MessageResponse, error) {
	var group model.Group
	db.First(&group, "uuid = ?", toUuid)
	if group.ID <= 0 {
		return nil, errors.New("群组不存在")
	}

	var messages []response.MessageResponse

	db.Raw("SELECT m.id, m.from_user_id, m.to_user_id, m.content, m.content_type, m.url, m.created_at, u.username AS from_username, u.avatar FROM messages AS m LEFT JOIN users AS u ON m.from_user_id = u.id WHERE m.message_type = 2 AND m.to_user_id = ?",
		group.ID).Scan(&messages)

	return messages, nil
}

func (m *messageService) SaveMessage(message protocol.Message) {
	db := pool.GetDB()
	var fromUser model.User
	db.Find(&fromUser, "uuid = ?", message.From)
	if NULL_ID == fromUser.Id {
		log.Logger.Error("SaveMessage not find from user", log.Any("SaveMessage not find from user", fromUser.Id))
		return
	}

	var toUserId int32 = 0

	if message.MessageType == constant.MESSAGE_TYPE_USER {
		var toUser model.User
		db.Find(&toUser, "uuid = ?", message.To)
		if NULL_ID == toUser.Id {
			return
		}
		toUserId = toUser.Id
	}

	if message.MessageType == constant.MESSAGE_TYPE_GROUP {
		var group model.Group
		db.Find(&group, "uuid = ?", message.To)
		if NULL_ID == group.ID {
			return
		}
		toUserId = group.ID
	}

	saveMessage := model.Message{
		FromUserId:  fromUser.Id,
		ToUserId:    toUserId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MessageType),
		Url:         message.Url,
	}
	db.Save(&saveMessage)
}
