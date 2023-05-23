package models

type Peer struct {
	Id          uint64 `json:"id"`
	Nickname    string `json:"nickname"`
	SchoolEmail string `json:"school_email"`
	Tribe     uint64 `json:"tribe"`
}
