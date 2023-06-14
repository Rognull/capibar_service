package handlers

import (
	"capi_api/internals/app/processors"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type CandidateHandler struct {
	processor *processors.CandidateProcessor
}

func NewCandidateHandler(processor *processors.CandidateProcessor) *CandidateHandler {
	handler := new(CandidateHandler)
	handler.processor = processor
	return handler
}


func (handler *CandidateHandler) GetCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) 
	if vars["convocation_id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	convocation_request, err := strconv.ParseInt(vars["convocation_id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}
	user, err := handler.processor.FindCandidate(uint64(convocation_request))
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
