package db

import (
    "fmt"
    "log"
    "database/sql"
    "github.com/abhisek1985/go-mux-CRUD/models" // models package where DB table schema is defined
)


// get all teams from the DB
func GetAllTeams(intPageNum int, intPageSize int) ([]models.Team, error) {
    // create the postgres db connection
    db := createDBConnection()

    // Slice of type structure
    teams := []models.Team{}

    // create the select sql query
    sqlStatement := `SELECT * FROM team ORDER BY ID DESC`

    if intPageNum > 0 && intPageSize > 0{
        paginatedQuery := models.PaginateQuery(intPageNum, intPageSize, sqlStatement)
        sqlStatement = paginatedQuery
    }

    rows, err := db.Query(sqlStatement)
    if err != nil{
       return teams, err
    }

    // iterate over the rows
    for rows.Next() {
        var team models.Team
        // unmarshal the row object to team
        err := rows.Scan(&team.ID, &team.Email, &team.MerchantID)
        if err != nil {
            break
        }
        // append the team in the teams slice
        teams = append(teams, team)
    }

    // close the statement
    defer rows.Close()

    // return empty team on error
    return teams, err
}

// create one team in the DB
func InsertTeam(team models.Team) (error, int) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the insert sql query
    // returning id of the inserted team
    sqlStatement := `INSERT INTO team (email, merchant_id) VALUES ($1, $2) RETURNING id`
    // the inserted id will store in this id
    var id int
    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, team.Email, team.MerchantID).Scan(&id)
    if err != nil {
        return err, 0
    }else{
        return nil, id
    }
}


// get one team from the DB by its teamId
func GetTeam(teamId int) (models.Team, error) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create a merchant of models.Merchant type
    var team models.Team
    // create the select sql query
    sqlStatement := `SELECT * FROM team WHERE id = $1;`
    // execute the sql statement
    row := db.QueryRow(sqlStatement, teamId)
    // unmarshal the row object to merchant
    err := row.Scan(&team.ID, &team.Email, &team.MerchantID)
    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return team, nil
    case nil:
        return team, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }
    // return empty team on error
    return team, err
}


// update one team in the DB by its teamId
func UpdateTeam(teamId int, team models.Team) int {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the update sql query
    sqlStatement := `UPDATE team SET email=$2, merchant_id=$3 WHERE id=$1;`
    // execute the sql statement
    res, err := db.Exec(sqlStatement, teamId, team.Email, team.MerchantID)
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    // check how many rows affected
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }
    fmt.Printf("Total rows/record affected %v", rowsAffected)
    return int(rowsAffected)
}


// delete team from the DB by teamId
func DeleteTeam(teamId int) int {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the delete sql query
    sqlStatement := `DELETE FROM team WHERE id = $1;`
    // execute the sql statement
    res, err := db.Exec(sqlStatement, teamId)
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    // check how many rows affected
    rowsAffected, err := res.RowsAffected()
    if err != nil {
        log.Fatalf("Error while checking the affected rows. %v", err)
    }
    fmt.Printf("Total rows/record affected %v", rowsAffected)
    return int(rowsAffected)
}


// Get Teams related to given merchantId
func GetMerchantTeamMembers(merchantId int, intPageNum int, intPageSize int) ([]models.Team, error){
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    var teams []models.Team
    // create the delete sql query
    sqlStatement := `SELECT * FROM team WHERE merchant_id = $1`

    if intPageNum > 0 && intPageSize > 0{
        paginatedQuery := models.PaginateQuery(intPageNum, intPageSize, sqlStatement)
        sqlStatement = paginatedQuery
    }

    // execute the sql statement
    rows, err := db.Query(sqlStatement, merchantId)
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    // close the statement
    defer rows.Close()
    // iterate over the rows
    for rows.Next() {
        var team models.Team
        // unmarshal the row object to team
        err = rows.Scan(&team.ID, &team.Email, &team.MerchantID)
        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }
        // append the team in the teams slice
        teams = append(teams, team)
    }
    // return empty team on error
    return teams, err
}