package main

import (
	"apifnf/db"
	"apifnf/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
