package models

type Message struct {
	Id           int64  `json:"id"`
	Body         string `json:"body"`
	Sender       string `json:"sender"`
	Recipient    string `json:"recipient"`
	AttachedPath string `json:"attached_path"`
	AttachedName string `json:"attached_name"`
	Time         int64  `json:"time"`
}
