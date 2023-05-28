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

// func (processor *VoteProcessor) CreatePeer(Vote models.Peer) error {

// 	return processor.storage.FindPeer(Vote)
// }

// func (processor *PeerProcessor) CreatePeer(peer models.Peer) error { //Процессор выполняет внутреннюю бизнес логику и валидирует переменные в соотвествии с ней
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

// 	return processor.storage.CreatePeer(peer)
// }

func (processor *PeerProcessor) FindPeer(peer_nickname string) (models.Peer, error) {
	peer := processor.storage.FindPeerByNickname(peer_nickname)

	if peer.Nickname == "" {
		return peer, errors.New("nickname should not be empty")
	}

	if peer.SchoolEmail == "" {
		return peer, errors.New("Email should not be empty")
	}

	// peer.Tribe < 0 || peer.Tribe > 4 линтер отругался на то, что peer.Tribe это uint и нет смысла его проверять на меньше нуля
	if peer.Tribe > 4 {
		return peer, errors.New("ALKOZAVR TRIBE")
	} //обязательно должен быть указан бренд

	// if peer.Code < 1000 || peer.Code > 9999 {
	// 	return  peer, errors.New("Code should not be empty")
	// }

	return peer, nil
}

// func (processor *PeerProcessor) ListPeers(peer_Nick string, Code string) ([]models.Peer, error) {
// 	return processor.storage.GetPeerList(peer_Nick, peer_Nick), nil
// }
