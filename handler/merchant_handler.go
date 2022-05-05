package handler

import (
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "log"
    "net/http" // used to access the request and response object of the api
    "strconv"  // package used to covert string into int type
    "github.com/abhisek1985/go-mux-CRUD/models" // models package where DB table schema is defined
    "github.com/abhisek1985/go-mux-CRUD/db" // db package where DB access handlers are present for API
    "github.com/gorilla/mux" // used to get the params from the route
)

// response format
type response struct {
    ID      int  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

// CreateMerchant create a merchant in the postgres db
func CreateMerchant(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
    // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // create an empty merchant of type models.Merchant
    var merchant models.Merchant
    // decode the json request to merchant
    err := json.NewDecoder(r.Body).Decode(&merchant)
    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
    // call insert merchant function and pass the merchant
    message := db.InsertMerchant(merchant)
    // format a response object
    res := response{
        Message: message,
    }
    // send the response
    json.NewEncoder(w).Encode(res)
}

// GetMerchant will return a single merchant by its id
func GetMerchant(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
   // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the merchantID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }
    // call the getUser function with user id to retrieve a single user
    user, err := db.GetMerchant(id)
    if err != nil {
        log.Fatalf("Unable to get merchant. %v", err)
    }
    // send the response
    json.NewEncoder(w).Encode(user)
}

// GetAllMerchant will return all the merchants
func GetAllMerchant(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
   // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get all the users in the db
    merchants, err := db.GetAllMerchants()
    if err != nil {
        log.Fatalf("Unable to get all merchants. %v", err)
    }
    // send all the merchants as response
    json.NewEncoder(w).Encode(merchants)
}

// UpdateMerchant update merchant's detail in the postgres db
func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
   // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the merchantID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }
    // create an empty merchant of type models.Merchant
    var merchant models.Merchant
    // decode the json request to merchant
    err = json.NewDecoder(r.Body).Decode(&merchant)
    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
    // call update merchant to update the merchant
    updatedRows := db.UpdateMerchant(id, merchant)
    // format the message string
    msg := fmt.Sprintf("Merchant updated successfully. Total rows/record affected %v", updatedRows)
    // format the response message
    res := response{
        ID:      id,
        Message: msg,
    }
    // send the response
    json.NewEncoder(w).Encode(res)
}

// DeleteMerchant delete merchant's detail in the postgres db
func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
    // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the merchantID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the id in string to int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }
    // call the DeleteMerchant, convert the int to int64
    deletedRows := db.DeleteMerchant(id)
    var message string
    if deletedRows == 0{
        message = fmt.Sprintf("Merchant deletion unsuccessful. Total rows/record affected %v", deletedRows)
    }else{
        message = fmt.Sprintf("Merchant deletion successful. Total rows/record affected %v", deletedRows)
    }
    // format the response message
    res := response{
        Message: message,
    }
    // send the response
    json.NewEncoder(w).Encode(res)
}

