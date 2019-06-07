package main

import (
	"cran/pkg/generating"
	_ "cran/pkg/generating/assembleenationale"
	"cran/pkg/http"
	"log"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		port = 8080
	}

	s := http.NewServer(generating.NewService())

	log.Fatal(s.Start(port))
}
