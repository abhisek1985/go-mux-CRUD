package test

import (
    "log"
    "testing"
    "bytes"
    "net/http"
    "net/http/httptest"
    "encoding/json"
)

func ensureMerchantTableExists() {
    db := getDBConnection()
	if _, err := db.Exec(merchantTableCreationQuery); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}


func clearMerchantTable() {
    db := getDBConnection()
    defer db.Close()
	db.Exec("DELETE FROM merchant")
}


func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	muxRouter := getMUXRouter()
	muxRouter.ServeHTTP(rr, req)
	return rr
}


func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestEmptyMerchants(t *testing.T) {
    ensureMerchantTableExists()
	clearMerchantTable()
	req, _ := http.NewRequest("GET", "/api/merchants", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
    responseBody := response.Body.String()
	if responseBody != "[]" {
    		t.Errorf("Expected an empty array. Got %s", responseBody)
    }
}

// TestCreateProduct Create Merchant
func TestCreateMerchant(t *testing.T) {
    ensureMerchantTableExists()
	clearMerchantTable()

	var jsonStr = []byte(`{"name":"test merchant", "code": "TEST-MER-1234"}`)
	req, _ := http.NewRequest("POST", "/api/create/merchant", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var merchant map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &merchant)

	if merchant["name"] != "test merchant" {
		t.Errorf("Expected merchant name to be 'test merchant'. Got '%v'", merchant["name"])
	}

	if merchant["code"] != "TEST-MER-1234" {
		t.Errorf("Expected merchant code to be 'TEST-MER-1234'. Got '%v'", merchant["code"])
	}
}


// TestGetNonExistentMerchant Fetch Non Exist Merchant
func TestGetNonExistentMerchant(t *testing.T) {
    ensureMerchantTableExists()
	clearMerchantTable()

	req, _ := http.NewRequest("GET", "/api/merchant/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Merchant not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Merchant not found'. Got '%s' ", m["error"])
	}
}

// TestUpdateNonExistentMerchant update Non Exist Merchant
func TestUpdateNonExistentMerchant(t *testing.T) {
    ensureMerchantTableExists()
	clearMerchantTable()

    var jsonStr = []byte(`{"name":"test merchant 11", "code": "TEST-MER-11"}`)
	req, _ := http.NewRequest("PUT", "/api/update/merchant/11", bytes.NewBuffer(jsonStr))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "No merchant found with id 11" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Merchant not found'. Got '%s' ", m["error"])
	}
}

// TestUpdateNonExistentMerchant update Non Exist Merchant
func TestDeleteNonExistentMerchant(t *testing.T) {
    ensureMerchantTableExists()
	clearMerchantTable()

	req, _ := http.NewRequest("DELETE", "/api/delete/merchant/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "No merchant found with id 11" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Merchant not found'. Got '%s' ", m["error"])
	}
}