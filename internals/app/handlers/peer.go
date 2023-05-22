package handlers

import (
	// "capi_api/internals/app/models"
	"capi_api/internals/app/processors"
	// "encoding/json"
	// "io"

	"errors"
	"github.com/gorilla/mux"
	"net/http"
	// "strconv"
	// "strings"
)

type PeerHandler struct {
	processor *processors.PeerProcessor
}

func NewPeerHandler(processor *processors.PeerProcessor) *PeerHandler {
	handler := new(PeerHandler)
	handler.processor = processor
	return handler
}



// func (handler *PeerHandler) FindPeer(w http.ResponseWriter, r *http.Request) {
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
	
// 	err = handler.processor.FindPeer(newVote)
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


func (handler *PeerHandler) GetPeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //переменные, обьявленные в ресурсах попадают в Vars и могут быть адресованы
	if vars["nickname"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	// peer_request, err := strconv.ParseInt(vars["id"], 10, 64)
	// if err != nil {
	// 	WrapError(w, err)
	// 	return
	// }

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
