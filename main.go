package main

import (
	"log"
	"os"
)

func main() {
	archivo, err := os.OpenFile("logs.log", os.O_CREATE|os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer archivo.Close()

	log.SetOutput(archivo)

	for i := 0; i < 10; i++ {
		log.Printf("Error linea %v\n", i)
	}

}
