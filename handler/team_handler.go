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


// CreateTeam create a team in the postgres db
func CreateTeam(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
    // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // create an empty team of type models.Team
    var team models.Team
    // decode the json request to team
    err := json.NewDecoder(r.Body).Decode(&team)
    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
    // call insert merchant function and pass the merchant
    message := db.InsertTeam(team)
    // format a response object
    res := response{
        Message: message,
    }
    // send the response
    json.NewEncoder(w).Encode(res)
}


// GetAllTeam will return all the teams
func GetAllTeam(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
    // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get all the users in the db
    teams, err := db.GetAllTeams()
    if err != nil {
        log.Fatalf("Unable to get all teams. %v", err)
    }
    // send all the merchants as response
    json.NewEncoder(w).Encode(teams)
}


// GetTeam will return a single team by its id
func GetTeam(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
   // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the teamID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }
    // call the GetTeam function with teamID to retrieve a single team
    user, err := db.GetTeam(id)
    if err != nil {
        log.Fatalf("Unable to get team. %v", err)
    }
    // send the response
    json.NewEncoder(w).Encode(user)
}


// UpdateTeam update team's detail in the postgres db
func UpdateTeam(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
   // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the teamID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }
    // create an empty team of type models.Team
    var team models.Team
    // decode the json request to team
    err = json.NewDecoder(r.Body).Decode(&team)
    if err != nil {
        log.Fatalf("Unable to decode the request body.  %v", err)
    }
    // call update UpdateTeam to update the merchant
    updatedRows := db.UpdateTeam(id, team)
    // format the message string
    msg := fmt.Sprintf("Team updated successfully. Total rows/record affected %v", updatedRows)
    // format the response message
    res := response{
        ID:      id,
        Message: msg,
    }
    // send the response
    json.NewEncoder(w).Encode(res)
}


// DeleteTeam delete team's detail in the postgres db
func DeleteTeam(w http.ResponseWriter, r *http.Request) {
    // set response header content type as application/json
    // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the teamID from the request params, key is "id"
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
        w.WriteHeader(http.StatusOK)
        message = fmt.Sprintf("Team deletion unsuccessful. Total rows/record affected %v", deletedRows)
    }else{
        w.WriteHeader(http.StatusOK)
        message = fmt.Sprintf("Team deletion successful. Total rows/record affected %v", deletedRows)
    }
    // format the response message
    res := response{
        Message: message,
    }
    // send the response
    json.NewEncoder(w).Encode(res)
}


// GetTeamsForMerchant return Teams related to given merchantID
func GetTeamsForMerchant(w http.ResponseWriter, r *http.Request){
    // set response header content type as application/json
    // Allow all origin to handle Cross-Origin Resource Sharing (CORS) issue
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // get the merchantID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the merchant_id in string to int
    merchantId, err := strconv.Atoi(params["merchant_id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }
    // get all the teams related to merchant from the db
    teams, err := db.GetMerchantTeamMembers(merchantId)
    if err != nil {
        log.Fatalf("Unable to get all teams. %v", err)
    }
    if len(teams) > 0{
        w.WriteHeader(http.StatusOK)
        // send all the merchants as response
        json.NewEncoder(w).Encode(teams)
    }else{
        s := make([]string, 0)
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(s)
    }
}