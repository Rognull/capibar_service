package handlers

import (
	"capi_api/internals/app/models"
	"capi_api/internals/app/processors"
	"encoding/json"
	"net/http"
)

type VoteHandler struct {
	processor *processors.VoteProcessor
}

func NewVoteHandler(processor *processors.VoteProcessor) *VoteHandler {
	handler := new(VoteHandler)
	handler.processor = processor
	return handler
}



func (handler *VoteHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var newVote models.Vote
	
	err := json.NewDecoder(r.Body).Decode(&newVote)
	if err != nil {
		WrapError(w, err)
		return
	}
	
	err = handler.processor.CreateVote(newVote)
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
