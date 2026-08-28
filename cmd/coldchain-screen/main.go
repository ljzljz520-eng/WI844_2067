package main

import (
	"coldchain/internal/coldchain"
	"coldchain/internal/httpapi"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	path := "coldchain.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	a, e := coldchain.NewApp(path)
	if e != nil {
		log.Fatal(e)
	}
	defer a.Close()
	fmt.Println("cold-chain dispatch screen listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", httpapi.New(a.S)))
}
