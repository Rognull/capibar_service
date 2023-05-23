package models

type Peer struct {
	Id          uint64 `json:"id"`
	Nickname    string `json:"nickname"`
	SchoolEmail string `json:"schoolemail"`
	Tribe     uint64 `json:"tribe"`
	Code        uint64 `json:"code"`
}
