// source: https://medium.com/towardsdev/golang-solid-principles-e7641dee90b0
package main

import (
	"fmt"
)

type DBConn interface {
	QuerySomeData() []string
}

// Defines the database connection
type MySQL struct{}

func (db MySQL) QuerySomeData() []string {
	return []string{"inf1", "inf2", "inf3"}
}

type MyRepository struct {
	db DBConn
}

func (r MyRepository) GetData() []string {
	return r.db.QuerySomeData()
}

func main() {
	mysqlDB := MySQL{}
	repo := MyRepository{db: mysqlDB}
	fmt.Println(repo.GetData())
}
