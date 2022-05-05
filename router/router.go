package router

import (
    "github.com/abhisek1985/go-mux-CRUD/handler"

    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/merchant", handler.GetAllMerchant).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/merchant/{id}", handler.GetMerchant).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/new-merchant", handler.CreateMerchant).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/merchant/{id}", handler.UpdateMerchant).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/del-merchant/{id}", handler.DeleteMerchant).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/team", handler.GetAllTeam).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/new-team", handler.CreateTeam).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/team/{id}", handler.GetTeam).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/team/{id}", handler.UpdateTeam).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/del-team/{id}", handler.DeleteTeam).Methods("DELETE", "OPTIONS")
    router.HandleFunc("/api/team/merchant/{merchant_id}", handler.GetTeamsForMerchant).Methods("GET", "OPTIONS")
    return router
}