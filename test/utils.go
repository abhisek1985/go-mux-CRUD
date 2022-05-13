package test

import (
    "fmt"
    "database/sql"
    "github.com/abhisek1985/go-mux-CRUD/handler"
    _ "github.com/lib/pq"      // postgres golang driver
    "github.com/gorilla/mux"
)


func getDBConnection() *sql.DB {
    var username string = APP_DB_USERNAME
    var password string = APP_DB_PASSWORD
    var database string = APP_DB_NAME
    var host string = APP_DB_HOST
    var port int = APP_DB_PORT


    // db_conn_str := fmt.Sprintf("user=%s dbname=%s sslmode=disable", username, database)
    db_conn_str := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, username, password, database)

    // Open the connection
    db, err := sql.Open("postgres", db_conn_str)
    if err != nil {
        panic(err)
    }
    // check the connection
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    // return the connection
    return db
}

func getMUXRouter() *mux.Router {
     router := mux.NewRouter()
     // APIs for Merchant
     router.HandleFunc("/api/merchants", handler.GetAllMerchant).Methods("GET", "OPTIONS")
     router.HandleFunc("/api/create/merchant", handler.CreateMerchant).Methods("POST", "OPTIONS")
     router.HandleFunc("/api/merchant/{id}", handler.GetMerchant).Methods("GET", "OPTIONS")
     router.HandleFunc("/api/update/merchant/{id}", handler.UpdateMerchant).Methods("PUT", "OPTIONS")
     router.HandleFunc("/api/delete/merchant/{id}", handler.DeleteMerchant).Methods("DELETE", "OPTIONS")

     router.HandleFunc("/api/teams", handler.GetAllTeam).Methods("GET", "OPTIONS")
     router.HandleFunc("/api/create/team", handler.CreateTeam).Methods("POST", "OPTIONS")
     return router
}