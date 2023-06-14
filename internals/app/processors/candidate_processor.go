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

func (processor *CandidateProcessor) FindCandidate(Convocation uint64) ([]models.CandidateWithName, error) {
	candidate := processor.storage.GetCandidateList(Convocation)
	return candidate, nil
}
