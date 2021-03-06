package db

import (
    "github.com/abhisek1985/go-mux-CRUD/models" // models package where DB table schema is defined
)


// get all merchant from the DB
func GetAllMerchants(intPageNum int, intPageSize int) ([]models.Merchant, error) {
    // create the postgres db connection
    db := createDBConnection()

    // close the db connection
    defer db.Close()

    // Slice of type structure
    merchants := []models.Merchant{}

    // create the select sql query
    sqlStatement := `SELECT * FROM merchant ORDER BY ID DESC`

    if intPageNum > 0 && intPageSize > 0{
        paginatedQuery := models.PaginateQuery(intPageNum, intPageSize, sqlStatement)
        sqlStatement = paginatedQuery
    }

    // execute the sql statement
    rows, err := db.Query(sqlStatement)
    if err != nil{
        return merchants, err
    }

    // iterate over the rows
    for rows.Next() {
        // structure variable
        var merchant models.Merchant
        // unmarshal the row object to merchant
        err := rows.Scan(&merchant.ID, &merchant.Code, &merchant.Name)
        if err != nil {
            break
        }
        // append the merchant in the merchants slice
        merchants = append(merchants, merchant)
    }

    // close the statement
    defer rows.Close()

    // return empty merchant on error
    return merchants, err
}


// create one merchant in the DB
func InsertMerchant(merchant models.Merchant) (error, int) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()
    // create the insert sql query
    // returning id of the inserted merchant
    sqlStatement := `INSERT INTO merchant (code, name) VALUES ($1, $2) RETURNING id`
    // the inserted id will store in this id
    var id int
    // execute the sql statement
    // Scan function will save the data and insert id in the id
    err := db.QueryRow(sqlStatement, merchant.Code, merchant.Name).Scan(&id)
    if err != nil {
        return err, 0
    }else{
        return nil, id
    }
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
    return merchant, err
}

// update one merchant in the DB by its merchantId
func UpdateMerchant(merchantId int, merchant models.Merchant) (error, int64) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()

    var rowsAffected int64 = 0
    // create the update sql query
    sqlStatement := `UPDATE merchant SET code=$2, name=$3 WHERE id=$1;`
    // execute the sql statement
    res, err := db.Exec(sqlStatement, merchantId, merchant.Code, merchant.Name)
    if err != nil{
        return err, rowsAffected
    }

    // check how many rows affected
    rowsAffected, updateError := res.RowsAffected()
    return updateError, rowsAffected
}

// delete merchant from the DB by merchantId
func DeleteMerchant(merchantId int) (error, int64) {
    // create the postgres db connection
    db := createDBConnection()
    // close the db connection
    defer db.Close()

    // create the delete sql query
    sqlStatement := `DELETE FROM merchant WHERE id = $1;`
    var rowsDeleted int64 = 0

    // execute the sql statement
    res, err := db.Exec(sqlStatement, merchantId)
    if err != nil {
        return err, rowsDeleted
    }

    // check how many rows deleted
    rowsDeleted, deleteError := res.RowsAffected()
    return deleteError, rowsDeleted
}
