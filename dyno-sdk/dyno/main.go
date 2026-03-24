package main

import (
	"log"
	"net/http"
	"time"

	dynosdk "github.com/winnerx0/dyno-sdk"
)

func main() {

	client := dynosdk.NewClient("dyno_qlDBXH0tAtsUEtGBXZOmdNEqYEjytocJL8SvJou2DXo")

	http.HandleFunc("/", client.NewLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(time.Second * 1)
		w.WriteHeader(201)
		w.Write([]byte("Hey this is a post"))
		
	})))

	log.Println("listening to test server")
	err := http.ListenAndServe(":3005", nil)
	if err != nil {
		log.Fatal(err)
	}
}
