package main

import (
	"excel-read/db"
	"excel-read/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
