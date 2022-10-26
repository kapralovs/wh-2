package main

import (
	"log"

	"github.com/kapralovs/wh-2/internal/server"
)

func main() {
	s := server.New()

	log.Fatal(s.Run())
}
