package processors

import (
	"capi_api/internals/app/db"
	"capi_api/internals/app/models"
	"errors"
)

type PeerProcessor struct {
	storage *db.PeerStorage
}

func NewPeerProcessor(storage *db.PeerStorage) *PeerProcessor {
	processor := new(PeerProcessor)
	processor.storage = storage
	return processor
}

func (processor *PeerProcessor) FindPeer(peer_nickname string) (models.Peer, error) {
	peer := processor.storage.FindPeerByNickname(peer_nickname)

	if peer.Nickname == "" {
		return peer, errors.New("nickname should not be empty")
	}

	if peer.SchoolEmail == "" {
		return peer, errors.New("Email should not be empty")
	}

	if peer.Tribe > 4 {
		return peer, errors.New("ALKOZAVR TRIBE")
	} 


	return peer, nil
}
