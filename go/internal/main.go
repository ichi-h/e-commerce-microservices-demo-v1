package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

const initQuery = `
DROP TABLE IF EXISTS members;
CREATE TABLE members (
  id INTEGER NOT NULL PRIMARY KEY,
  name VARCHAE(20) DEFAULT ''
);
INSERT INTO members (name) VALUES ("Steve");
INSERT INTO members (name) VALUES ("Bob");
INSERT INTO members (name) VALUES ("Alice");`

type Member struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func main() {
	db, err := sqlx.Connect("sqlite3", "db/test.db")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(initQuery)

	chris := &Member{
		Name: "Chris",
	}
	transaction := db.MustBegin()
	transaction.NamedExec("INSERT INTO members (name) VALUES (:name)", chris)
	transaction.Commit()

	members := []Member{}
	err = db.Select(&members, "SELECT id, name FROM members")
	if err != nil {
		log.Fatalln(err)
	}

	for _, member := range members {
		fmt.Println(member.ID, member.Name)
	}

	e := echo.New()
	e.GET("/", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
