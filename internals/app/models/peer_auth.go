package models

type PeerAuth struct {
	PeerId   uint64 `json:"peer_id"`
	Password string `json:"password"`
}
