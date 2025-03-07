// In this below we are checking that which database has to be connect when user lookup something.
// User does not know from which database the output is coming from

package main

import "fmt"

type Database interface {
	Connect()
}

type MySql struct{}

func (m MySql) Connect() {
	fmt.Println("MySql is Connected")
}

type PostgreSQL struct{}

func (p PostgreSQL) Connect() {
	fmt.Println("PostgreSQL is connected")
}

func DatabaseFactory(db string) Database {
	if db == "mySQL" {
		return MySql{}
	} else if db == "postgreSQL" {
		return PostgreSQL{}
	}
	return nil
}

func main() {

	db := DatabaseFactory("mySQL")
	db.Connect()

	db2 := DatabaseFactory("postgreSQL")
	db2.Connect()

}
