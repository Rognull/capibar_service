package models

type Peer struct {
	Id          uint64 `json:"id"`
	NickName    string `json:"nickname"`
	SchoolEmail string `json:"email"`
	TribeId     uint64 `json:"tribe"`
	Code        uint64 `json:"code"`
}
