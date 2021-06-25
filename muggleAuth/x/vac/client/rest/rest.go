package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers vac-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/vac/credentials/{id}", getCredentialHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/vac/credentials", listCredentialHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/vac/credentials", createCredentialHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/vac/credentials/{id}", updateCredentialHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/vac/credentials/{id}", deleteCredentialHandler(clientCtx)).Methods("POST")

}
