package request

type MessageRequest struct {
	MessageType    int32  `json:"messageType"`
	Uuid           string `json:"uuid"`
	FriendUsername string `json:"friendUsername"`
}
