package processors

import (
	"capi_api/internals/app/db"
	"capi_api/internals/app/models"
 
)

type CandidateProcessor struct {
	storage *db.CandidateStorage
}

func NewCandidateProcessor(storage *db.CandidateStorage) *CandidateProcessor {
	processor := new(CandidateProcessor)
	processor.storage = storage
	return processor
}

// func (processor *VoteProcessor) CreateCandidate(Vote models.Candidate) error {

// 	return processor.storage.FindCandidate(Vote)
// }

// func (processor *CandidateProcessor) CreateCandidate(peer models.Candidate) error { //Процессор выполняет внутреннюю бизнес логику и валидирует переменные в соотвествии с ней
// 	if peer.NickName == "" {
// 		return errors.New("nickname should not be empty")
// 	}

// 	if peer.SchoolEmail == "" {
// 		return errors.New("Email should not be empty")
// 	}

// 	if peer.TribeId < 0 || peer.TribeId > 4 {
// 		return errors.New("ALKOZAVR TRIBE")
// 	} //обязательно должен быть указан бренд

// 	if peer.Code == "" {
// 		return errors.New("Code should not be empty")
// 	}

// 	return processor.storage.CreateCandidate(peer)
// }

func (processor *CandidateProcessor) FindCandidate(Convocation uint64) ([]models.CandidateWithName, error) {
	candidate := processor.storage.GetCandidateList(Convocation)
	return candidate, nil
}

// func (processor *CandidateProcessor) ListCandidates(peer_Nick string, Code string) ([]models.Candidate, error) {
// 	return processor.storage.GetCandidateList(peer_Nick, peer_Nick), nil
// }
