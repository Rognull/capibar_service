package handlers

import (
	"capi_api/internals/app/models"
	"capi_api/internals/app/processors"
	"encoding/json"
	"fmt"
	// "io"

	// "errors"
	// "github.com/gorilla/mux"
	"net/http"
	// "strconv"
	// "strings"
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

	// b, _ := io.ReadAll(r.Body)
	// // b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
	

	// fmt.Println(string(b))
	
	err := json.NewDecoder(r.Body).Decode(&newVote)
	if err != nil {
		// fmt.Printf("%s",json.NewDecoder(r.Body).Decode(&newVote))
		WrapError(w, err)
		return
	}
	fmt.Println(newVote)
	
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


// func (handler *VoteHandler) List(w http.ResponseWriter, r *http.Request) {
// 	vars := r.URL.Query()
// 	list, err := handler.processor.ListUsers(strings.Trim(vars.Get("name"), "\""))

// 	if err != nil {
// 		WrapError(w, err)
// 	}

// 	var m = map[string]interface{} {
// 		"result" : "OK",
// 		"data" : list,
// 	}

// 	WrapOK(w, m)
// }

// func (handler *VoteHandler) Find(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
// 	if vars["id"] == "" {
// 		WrapError(w, errors.New("missing id"))
// 		return
// 	}

// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	if err != nil {
// 		WrapError(w, err)
// 		return
// 	}

// 	user, err := handler.processor.FindUser(id)
// 	if err != nil {
// 		WrapError(w, err)
// 		return
// 	}

// 	var m = map[string]interface{} {
// 		"result" : "OK",
// 		"data" : user,
// 	}

// 	WrapOK(w, m)
// }
