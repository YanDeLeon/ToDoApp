package models

import (
	"bytes"
	"log"
	base "todoapp"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v5"
)

// DB is a connection to your database to be used
// throughout your application.

func init() {
	bf, err := base.Config.Find("database.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = pop.LoadFrom(bytes.NewReader(bf))
	if err != nil {
		log.Fatal(err)
	}
}

// DB returns the DB connection for the current environment.
func DB() *pop.Connection {
	c, err := pop.Connect(envy.Get("GO_ENV", "development"))
	if err != nil {
		log.Fatal(err)
	}

	return c
}
