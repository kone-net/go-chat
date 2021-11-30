package response

import "chat-room/internal/model"

type SearchResponse struct {
	User  model.User  `json:"user"`
	Group model.Group `json:"group"`
}
