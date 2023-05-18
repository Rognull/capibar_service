package models

type Vote struct {
	Id_voter   int64 `json:"id_voter"`
	Id_candidate int64 `json:"candidate"`
	Id_election int64 `json:"election"`
}
 