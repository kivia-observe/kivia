package main

import (
	"log"
	"net/http"
	"time"
	_ "time"

	dynosdk "github.com/winnerx0/dyno-sdk"
)

func main() {

	httpc := http.Client{}

	// start := time.Now()

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := httpc.Do(req)

	// endTime := time.Since(start)

	if err != nil {
		log.Fatal(err)
	}

	client := dynosdk.NewClient("dyno_NHLfFHXlOdZHmJVL8gecgPcU0q5FvUbOchM0tvqt2eM", req, res)

	http.HandleFunc("/", client.NewLog(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		time.Sleep(time.Second * 5)
		w.WriteHeader(500)
		w.Write([]byte("Hey this is a post"))
		
	})))

	log.Println("listening to test server")
	http.ListenAndServe(":3000", nil)

}
