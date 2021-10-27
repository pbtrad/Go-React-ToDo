package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pbtrad/golang-react-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting the sever on port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}
