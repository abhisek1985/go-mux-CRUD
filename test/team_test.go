package test

import (
    "log"
    "testing"
    "bytes"
    "net/http"
    // "net/http/httptest"
    "encoding/json"
)

func ensureTeamTableExists() {
    db := getDBConnection()
	if _, err := db.Exec(teamTableCreationQuery); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}


func clearTeamTable() {
    db := getDBConnection()
    defer db.Close()
	db.Exec("DELETE FROM team")
}

// TestEmptyTeams no teams
func TestEmptyTeams(t *testing.T) {
    ensureTeamTableExists()
	clearTeamTable()
	req, _ := http.NewRequest("GET", "/api/teams", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
    responseBody := response.Body.String()
	if responseBody != "[]" {
    		t.Errorf("Expected an empty array. Got %s", responseBody)
    }
}


// TestCreateProduct Create Merchant
func TestCreateTeam(t *testing.T) {
    ensureTeamTableExists()
    clearTeamTable()

	var jsonStr = []byte(`{"email":"test_merchant@example.com"}`)
	req, _ := http.NewRequest("POST", "/api/create/team", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var team map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &team)

	if team["email"] != "test_merchant@example.com" {
		t.Errorf("Expected team email to be 'test_merchant@example.com'. Got '%v'", team["email"])
	}

	if team["merchant_id"] != nil {
		t.Errorf("Expected team merchant code to be 'Null'. Got '%v'", team["merchant_id"])
	}
}