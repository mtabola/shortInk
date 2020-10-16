package globalVars

import (
	"../structures"
	"database/sql"
)

var Links = structures.LinksList{
	Links: make([]structures.Link, 0),
}

var DB *sql.DB

var Response = structures.Response{nil, nil}
