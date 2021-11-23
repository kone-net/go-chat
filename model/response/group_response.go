package response

import "time"

type GroupResponse struct {
	Uuid      string    `json:"uuid"`
	GroupId   int32     `json:"groupId"`
	CreatedAt time.Time `json:"createAt"`
	Name      string    `json:"name"`
	Notice    string    `json:"notice"`
}
