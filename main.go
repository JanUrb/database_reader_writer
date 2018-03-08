package main

import (
	"log"

	"github.com/JanUrb/database_reader_writer/database"
)

func main() {
	data, err := database.Read()
	if err != nil {
		log.Println("Error wile reading from database")
		return
	}

	log.Println("Received data: ", data)
}
