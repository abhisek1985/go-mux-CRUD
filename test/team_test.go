package test

import (
    "log"
    "testing"
    // "bytes"
    "net/http"
    // "net/http/httptest"
    // "encoding/json"
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