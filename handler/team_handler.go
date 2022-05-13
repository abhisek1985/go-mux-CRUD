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

    //Payload validation
    validErrs := team.ValidateTeam()
    if len(validErrs) > 0{
        err := map[string]interface{}{"validationError": validErrs}
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(err)
        return
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
    var PageNum, PageSize string
    var intPageNum, intPageSize int
    var err error

    // Pagination
    PageNum = r.URL.Query().Get("PageNum")
    PageSize = r.URL.Query().Get("PageSize")
    intPageNum = 0
    intPageSize = 0

    if PageNum != "" {
        intPageNum, err = strconv.Atoi(PageNum)
        if err != nil {
           respondWithError(w, http.StatusBadRequest, "Invalid PageNum value")
           return
        }
    }

    if PageSize != "" {
        intPageSize, err = strconv.Atoi(PageSize)
        if err != nil {
            respondWithError(w, http.StatusBadRequest, "Invalid PageSize value")
            return
        }
    }

    // get all the users in the db
    teams, err := db.GetAllTeams(intPageNum, intPageSize)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Unable to fetch teams")
    }

    respondWithJSON(w, http.StatusOK, teams)
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
    team, err := db.GetTeam(id)
    if err != nil {
        log.Fatalf("Unable to get team. %v", err)
    }
    // send the response
    json.NewEncoder(w).Encode(team)
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
    deletedRows := db.DeleteTeam(id)
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

    var PageNum, PageSize string
    var intPageNum, intPageSize int
    var err error
    var merchantId int

    // get the merchantID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the merchant_id in string to int
    merchantId, err = strconv.Atoi(params["merchant_id"])
    if err != nil {
        log.Fatalf("Unable to convert the string into int.  %v", err)
    }

    // Pagination
    PageNum = r.URL.Query().Get("PageNum")
    PageSize = r.URL.Query().Get("PageSize")
    intPageNum = 0
    intPageSize = 0

    if PageNum != "" {
        intPageNum, err = strconv.Atoi(PageNum)
        if err != nil {
           log.Fatalf("Unable to get all teams. %v", err)
           return
        }
    }

    if PageSize != "" {
        intPageSize, err = strconv.Atoi(PageSize)
        if err != nil {
            log.Fatalf("Unable to get all teams. %v", err)
            return
        }
    }

    // get all the teams related to merchant from the db
    teams, err := db.GetMerchantTeamMembers(merchantId, intPageNum, intPageSize)
    if err != nil {
        log.Fatalf("Unable to get all teams. %v", err)
    }
    if len(teams) > 0{
        w.WriteHeader(http.StatusOK)
        // send all the teams as response
        json.NewEncoder(w).Encode(teams)
    }else{
        s := make([]string, 0)
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(s)
    }
}