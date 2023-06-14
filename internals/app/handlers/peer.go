package handlers

import (
	"capi_api/internals/app/processors"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type PeerHandler struct {
	processor *processors.PeerProcessor
}

func NewPeerHandler(processor *processors.PeerProcessor) *PeerHandler {
	handler := new(PeerHandler)
	handler.processor = processor
	return handler
}


func (handler *PeerHandler) GetPeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	if vars["nickname"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	user, err := handler.processor.FindPeer(vars["nickname"])
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : user,
	}

	WrapOK(w, m)
}
