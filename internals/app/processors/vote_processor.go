package processors

import (
	"capi_api/internals/app/db"
	"capi_api/internals/app/models"
	"errors"
)

type VoteProcessor struct {
	storage *db.VoteStorage
}

func NewVoteProcessor(storage *db.VoteStorage) *VoteProcessor {
	processor := new(VoteProcessor)
	processor.storage = storage
	return processor
}


func (processor *VoteProcessor) CreateVote(Vote models.Vote) error {
	if Vote.Id_voter < 0  {  // TODO CREAT ALL CHECKS 
		return errors.New("name should not be empty")
	}

	return processor.storage.InsertVote(Vote)
}
