package models

import (
    "fmt"
)

func PaginateQuery(PageNumber int, PageSize int, SQLQuery string) string{
    var PageOffset int
    PageOffset = (PageNumber - 1) * PageSize
    PaginatedQuery := SQLQuery + fmt.Sprintf(" limit %d offset %d", PageSize, PageOffset)
    return PaginatedQuery
}