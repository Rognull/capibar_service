package processors

import (
	"capi_api/internals/app/db"
	"capi_api/internals/app/models"
	"errors"
)

type PeerAuthProcessor struct {
	storage *db.PeerAuthStorage
}

func NewPeerAuthProcessor(storage *db.PeerAuthStorage) *PeerAuthProcessor {
	processor := new(PeerAuthProcessor)
	processor.storage = storage
	return processor
}


func (processor *PeerAuthProcessor) CreatePeerAuth(PeerAuth models.PeerAuth) error {
	if PeerAuth.PeerId < 1  {  // TODO CREAT ALL CHECKS 
		return errors.New("Id should not be empty")
	}

	return processor.storage.UpdatePassword(PeerAuth)
}