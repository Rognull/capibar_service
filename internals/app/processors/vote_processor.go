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

// func (processor *VoteProcessor) FindVote(id int64) (models.Vote, error) {
// 	Vote := processor.storage.GetVoteById(id)

// 	if Vote.Id != id {
// 		return Vote, errors.New("Vote not found")
// 	}

// 	return Vote, nil

// }

// func (processor *VoteProcessor) ListVotes(nameFilter string) ([]models.Vote, error) {
// 	return processor.storage.GetVotesList(nameFilter), nil
// }

