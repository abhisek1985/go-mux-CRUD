package db

import (
    "fmt"
    "log"
    "database/sql"
    "github.com/abhisek1985/go-mux-CRUD/models" // models package where DB table schema is defined
)


// get all merchant from the DB
func GetAllMerchants(intPageNum int, intPageSize int) ([]models.Merchant, error) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    var merchants []models.Merchant

    // create the select sql query
    sqlStatement := `SELECT * FROM merchant ORDER BY ID DESC`

    if intPageNum > 0 && intPageSize > 0{
        paginatedQuery := models.PaginateQuery(intPageNum, intPageSize, sqlStatement)
        sqlStatement = paginatedQuery
    }

    // execute the sql statement
    rows, err := db.Query(sqlStatement)
    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }
    // close the statement
    defer rows.Close()
    // iterate over the rows
    for rows.Next() {
        var merchant models.Merchant
        // unmarshal the row object to merchant
        err = rows.Scan(&merchant.ID, &merchant.Code, &merchant.Name)
        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }
        // append the merchant in the merchants slice
        merchants = append(merchants, merchant)
    }
    // return empty merchant on error
    return merchants, err
}


// create one merchant in the DB
func InsertMerchant(merchant models.Merchant) string {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the insert sql query
    // returning id of the inserted merchant
    sqlStatement := `INSERT INTO merchant (code, name) VALUES ($1, $2) RETURNING id`
    // the inserted id will store in this id
    var id int
    var message string
    // execute the sql statement
    // Scan function will save the insert id in the id
    err := db.QueryRow(sqlStatement, merchant.Code, merchant.Name).Scan(&id)
    if err != nil {
        message = fmt.Sprintf("Merchant creation unsuccessful reason %v", err)
    }else{
        fmt.Printf("Inserted a single record %v", id)
        message = "Merchant created successfully"
    }
    return message
}


// get one merchant from the DB by its merchantId
func GetMerchant(merchantId int) (models.Merchant, error) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create a merchant of models.Merchant type
    var merchant models.Merchant
    // create the select sql query
    sqlStatement := `SELECT * FROM merchant WHERE id = $1;`
    // execute the sql statement
    row := db.QueryRow(sqlStatement, merchantId)
    // unmarshal the row object to merchant
    err := row.Scan(&merchant.ID, &merchant.Code, &merchant.Name)
    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return merchant, nil
    case nil:
        return merchant, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }
    // return empty merchant on error
    return merchant, err
}

// update one merchant in the DB by its merchantId
func UpdateMerchant(merchantId int, merchant models.Merchant) int {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the update sql query
    sqlStatement := `UPDATE merchant SET code=$2, name=$3 WHERE id=$1;`
    // execute the sql statement
    res, err := db.Exec(sqlStatement, merchantId, merchant.Code, merchant.Name)
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

// delete merchant from the DB by merchantId
func DeleteMerchant(merchantId int) int {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the delete sql query
    sqlStatement := `DELETE FROM merchant WHERE id = $1;`
    // execute the sql statement
    res, err := db.Exec(sqlStatement, merchantId)
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
