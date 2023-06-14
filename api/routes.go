package api

import (
	"capi_api/internals/app/handlers"
	"github.com/gorilla/mux"
)

func CreateRoutes(voteHandler *handlers.VoteHandler, peerHandler *handlers.PeerHandler, peerAuthHandler *handlers.PeerAuthHandler,candidateHandler *handlers.CandidateHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/vote/insert", voteHandler.Insert).Methods("POST") 
	r.HandleFunc("/peer/find/{nickname}", peerHandler.GetPeer).Methods("GET")
	r.HandleFunc("/peer_auth", peerAuthHandler.Update).Methods("PATCH")
	r.HandleFunc("/convocation/candidates/{convocation_id}", candidateHandler.GetCandidate).Methods("GET")
	r.NotFoundHandler = r.NewRoute().HandlerFunc(handlers.NotFound).GetHandler()
	return r
}
