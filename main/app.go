package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"enigma-mhs/config"
	"enigma-mhs/delivery"
)

type app struct {
	db *sql.DB
}

func newApp() app {
	c := config.NewConfig()
	err := c.InitDB()
	if err != nil {
		panic(err)
	}
	myapp := app{
		db: c.Db,
	}
	return myapp
}

func (a app) cli() {
	a.run(delivery.NewAppDelivery(a.db))
}

func (a app) run(delivery *delivery.AppDelivery) {
	log.Println("Success")
	delivery.Run()
}

func main() {
	newApp().cli()
}


