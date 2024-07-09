package main

import (
	"fmt"
	"temperatureblanket/db"
)

func main() {

	conn, err := db.Init()
	if err != nil {
		fmt.Println("erro aqui", err)
	}

	fmt.Println(conn)

}
