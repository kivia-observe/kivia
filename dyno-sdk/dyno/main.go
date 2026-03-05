package main

import (
	"log"
	"net/http"
	"time"

	dynosdk "github.com/winnerx0/dyno-sdk"
)

func main() {

	httpc := http.Client{}

	start := time.Now()

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/todos/1", nil)

	if err != nil {
		log.Fatal(err)
	}

	res, err := httpc.Do(req)

	endTime := time.Since(start)

	if err != nil {
		log.Fatal(err)
	}

	client := dynosdk.NewClient("dyno_LRBG46d1I3FoPZrfH0nNI0iNmwWCANm7bfGK8Scbiw0", req, res)

	err = client.NewLog(endTime)

	if err != nil {
		log.Fatal(err)
	}

}
