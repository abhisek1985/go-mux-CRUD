package handler

import (
    "encoding/json" // package to encode and decode the json into struct and vice versa
    "fmt"
    "database/sql"
    "net/http" // used to access the request and response object of the api
    "strconv"  // package used to covert string into int type
    "github.com/abhisek1985/go-mux-CRUD/models" // models package where DB table schema is defined
    "github.com/abhisek1985/go-mux-CRUD/db" // db package where DB access handlers are present for API
    "github.com/gorilla/mux" // used to get the params from the route
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// CreateMerchant create a merchant in the postgres db
func CreateMerchant(w http.ResponseWriter, r *http.Request) {

    // create an empty merchant of type models.Merchant
    var merchant models.Merchant

    // decode the json request to merchant
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&merchant)

    //Unable to decode Payload
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Invalid request payload")
        return
    }

    defer r.Body.Close()

    //Payload validation
    if validErrs := merchant.ValidateMerchant(); len(validErrs) > 0{
        respondWithJSON(w, http.StatusBadRequest, validErrs)
        return
    }

    // call insert merchant function and pass the merchant
    err, id := db.InsertMerchant(merchant)

    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    merchant.ID = id
    respondWithJSON(w, http.StatusCreated, merchant)
}

// GetMerchant will return a single merchant by its id
func GetMerchant(w http.ResponseWriter, r *http.Request) {
    // get the merchantID from the request params, key is "id"
    params := mux.Vars(r)
    // convert the id type from string to int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Invalid query params type")
        return
    }
    // call the getUser function with user id to retrieve a single user
    merchant, err := db.GetMerchant(id)
    if err != nil{
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Merchant not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }
    respondWithJSON(w, http.StatusOK, merchant)
}

// GetAllMerchant will return all the merchants
func GetAllMerchant(w http.ResponseWriter, r *http.Request) {
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
    merchants, err := db.GetAllMerchants(intPageNum, intPageSize)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Unable to fetch merchants")
    }

    respondWithJSON(w, http.StatusOK, merchants)
}

// UpdateMerchant update merchant's detail in the postgres db
func UpdateMerchant(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    // Get merchant id from query param and convert it into int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid query params type")
        return
    }

    // create an empty merchant of type models.Merchant
    var merchant models.Merchant
    // decode the json request to merchant
    decoder := json.NewDecoder(r.Body)
    decodeError := decoder.Decode(&merchant)
    if decodeError != nil {
        respondWithError(w, http.StatusInternalServerError, "Invalid request payload")
        return
    }

    defer r.Body.Close()


    // call update merchant to update the merchant
    updateError, rowsAffected := db.UpdateMerchant(id, merchant)

    if updateError != nil{
        respondWithError(w, http.StatusInternalServerError, updateError.Error())
        return
    }

    if rowsAffected == 0{
        errorResp := fmt.Sprintf("No merchant found with id %d", id)
        respondWithError(w, http.StatusNotFound, errorResp)
        return
    }
    merchant.ID = id
    respondWithJSON(w, http.StatusCreated, merchant)
}

// DeleteMerchant delete merchant's detail in the postgres db
func DeleteMerchant(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    // Get merchant id from query param and convert it into int
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid query params type")
        return
    }

    // call the DeleteMerchant, convert the int to int64
    deleteError, deletedRows := db.DeleteMerchant(id)

    if deleteError != nil{
        respondWithError(w, http.StatusInternalServerError, deleteError.Error())
        return
    }

    if deletedRows == 0{
        errorResp := fmt.Sprintf("No merchant found with id %d", id)
        respondWithError(w, http.StatusNotFound, errorResp)
        return
    }

    respondWithJSON(w, http.StatusCreated, fmt.Sprintf("Merchant deleted with id %d successfully", id))
}

