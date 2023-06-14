package handlers

import (
	"capi_api/internals/app/models"
	"capi_api/internals/app/processors"
	"encoding/json"
	"net/http"
)

type PeerAuthHandler struct {
	processor *processors.PeerAuthProcessor
}

func NewPeerAuthHandler(processor *processors.PeerAuthProcessor) *PeerAuthHandler {
	handler := new(PeerAuthHandler)
	handler.processor = processor
	return handler
}



func (handler *PeerAuthHandler) Update(w http.ResponseWriter, r *http.Request) {
	var newPeerAuth models.PeerAuth
	
	err := json.NewDecoder(r.Body).Decode(&newPeerAuth)
	if err != nil {
		WrapError(w, err)
		return
	}
	
	err = handler.processor.CreatePeerAuth(newPeerAuth)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : "",
	}

	WrapOK(w, m)

}