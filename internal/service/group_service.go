package service

import (
	"chat-room/internal/dao/pool"
	"chat-room/pkg/common/response"
	"chat-room/pkg/errors"

	"chat-room/internal/model"

	"github.com/google/uuid"
)

type groupService struct {
}

var GroupService = new(groupService)

func (g *groupService) GetGroups(uuid string) ([]response.GroupResponse, error) {
	db := pool.GetDB()

	migrate := &model.Group{}
	pool.GetDB().AutoMigrate(&migrate)
	migrate2 := &model.GroupMember{}
	pool.GetDB().AutoMigrate(&migrate2)

	var queryUser *model.User
	db.First(&queryUser, "uuid = ?", uuid)

	if queryUser.Id <= 0 {
		return nil, errors.New("用户不存在")
	}

	var groups []response.GroupResponse

	db.Raw("SELECT g.id AS group_id, g.uuid, g.created_at, g.name, g.notice FROM group_members AS gm LEFT JOIN `groups` AS g ON gm.group_id = g.id WHERE gm.user_id = ?",
		queryUser.Id).Scan(&groups)

	return groups, nil
}

func (g *groupService) SaveGroup(userUuid string, group model.Group) {
	db := pool.GetDB()
	var fromUser model.User
	db.Find(&fromUser, "uuid = ?", userUuid)
	if fromUser.Id <= 0 {
		return
	}

	group.UserId = fromUser.Id
	group.Uuid = uuid.New().String()
	db.Save(&group)

	groupMember := model.GroupMember{
		UserId:   fromUser.Id,
		GroupId:  group.ID,
		Nickname: fromUser.Username,
		Mute:     0,
	}
	db.Save(&groupMember)
}

func (g *groupService) GetUserIdByGroupUuid(groupUuid string) []model.User {
	var group model.Group
	db := pool.GetDB()
	db.First(&group, "uuid = ?", groupUuid)
	if group.ID <= 0 {
		return nil
	}

	var users []model.User
	db.Raw("SELECT u.uuid, u.avatar, u.username FROM `groups` AS g JOIN group_members AS gm ON gm.group_id = g.id JOIN users AS u ON u.id = gm.user_id WHERE g.id = ?",
		group.ID).Scan(&users)
	return users
}

func (g *groupService) JoinGroup(groupUuid, userUuid string) error {
	var user model.User
	db := pool.GetDB()
	db.First(&user, "uuid = ?", userUuid)
	if user.Id <= 0 {
		return errors.New("用户不存在")
	}

	var group model.Group
	db.First(&group, "uuid = ?", groupUuid)
	if user.Id <= 0 {
		return errors.New("群组不存在")
	}
	var groupMember model.GroupMember
	db.First(&groupMember, "user_id = ? and group_id = ?", user.Id, group.ID)
	if groupMember.ID > 0 {
		return errors.New("已经加入该群组")
	}
	nickname := user.Nickname
	if nickname == "" {
		nickname = user.Username
	}
	groupMemberInsert := model.GroupMember{
		UserId:   user.Id,
		GroupId:  group.ID,
		Nickname: nickname,
		Mute:     0,
	}
	db.Save(&groupMemberInsert)

	return nil
}
