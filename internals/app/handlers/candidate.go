package handlers

import (
	// "capi_api/internals/app/models"
	"capi_api/internals/app/processors"
	// "encoding/json"
	// "io"
	// "fmt"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	// "strings"
)

type CandidateHandler struct {
	processor *processors.CandidateProcessor
}

func NewCandidateHandler(processor *processors.CandidateProcessor) *CandidateHandler {
	handler := new(CandidateHandler)
	handler.processor = processor
	return handler
}



// func (handler *CandidateHandler) FindCandidate(w http.ResponseWriter, r *http.Request) {
// 	var newVote models.Vote

// 	// b, _ := io.ReadAll(r.Body)
// 	// // b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	

// 	// fmt.Println(string(b))
	
// 	// err := json.NewDecoder(r.Body).Decode(&newVote)
// 	// if err != nil {
// 	// 	// fmt.Printf("%s",json.NewDecoder(r.Body).Decode(&newVote))
// 	// 	WrapError(w, err)
// 	// 	return
// 	// }
	
// 	err = handler.processor.FindCandidate(newVote)
// 	if err != nil {
// 		WrapError(w, err)
// 		return
// 	}

// 	var m = map[string]interface{} {
// 		"result" : "OK",
// 		"data" : "",
// 	}

// 	WrapOK(w, m)

// }


func (handler *CandidateHandler) GetCandidate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
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
