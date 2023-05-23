package models

type CandidateWithName struct {
	PeerId    uint64 `json:"peer_id"`
	Nickname string `json:"nickname"`
	ConvocationId  uint64 `json:"convocation_id"`
	Promises     string `json:"promises"`
}
