package router

import (
    "net/http"
    "github.com/abhisek1985/go-mux-CRUD/handler"
    "github.com/abhisek1985/go-mux-CRUD/middlewares"
    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    // APIs for Merchant
    router.HandleFunc("/api/merchants", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.GetAllMerchant))).Methods("GET", "OPTIONS")

    router.HandleFunc("/api/merchant/{id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.GetMerchant))).Methods("GET", "OPTIONS")

    router.HandleFunc("/api/create/merchant", handler.CreateMerchant).Methods("POST", "OPTIONS")

    router.HandleFunc("/api/update/merchant/{id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.UpdateMerchant))).Methods("PUT", "OPTIONS")

    router.HandleFunc("/api/delete/merchant/{id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.DeleteMerchant))).Methods("DELETE", "OPTIONS")

    // APIs for Team
    router.HandleFunc("/api/teams", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.GetAllTeam))).Methods("GET", "OPTIONS")

    router.HandleFunc("/api/team/{id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.GetTeam))).Methods("GET", "OPTIONS")

    router.HandleFunc("/api/create/team", handler.CreateTeam).Methods("POST", "OPTIONS")

    router.HandleFunc("/api/update/team/{id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.UpdateTeam))).Methods("PUT", "OPTIONS")

    router.HandleFunc("/api/delete/team/{id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.DeleteTeam))).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/teams/merchant/{merchant_id}", middlewares.BasicAuthMiddleware(http.HandlerFunc(
    handler.GetTeamsForMerchant))).Methods("GET", "OPTIONS")

    return router
}