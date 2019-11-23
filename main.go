package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/eezhal92/high-traffic/domain"
	"github.com/eezhal92/high-traffic/services"
)

// fetch two remote data
// both of them should be done under 2 seconds
// when one of them exceeding the limit, then cancel and return empty response
// otherwise return response
func main() {
	rand.Seed(time.Now().UnixNano())

	resp, err := getData()

	if err != nil {
		fmt.Println("err!! ", err)
	} else {
		fmt.Println("resp: ", resp)
	}
}

func getData() ([]string, error) {
	stopChan := make(chan bool, 1)

	c := make(chan string, 2)
	defer close(c)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var resp []string

	go fetchRemoteData(ctx, c, stopChan, "shipping")
	go fetchRemoteData(ctx, c, stopChan, "transaction")

	for i := 0; i < 2; i++ {
		select {
		case <-stopChan:
			fmt.Println("stopped")
			errorResp := make([]string, 0)
			return errorResp, fmt.Errorf("too long fetching data")
		case msg := <-c:
			resp = append(resp, msg)
		}
	}

	return resp, nil
}

func fetchRemoteData(ctx context.Context, c chan string, stopChan chan bool, successText string) {
	n := rand.Intn(4000)
	timeout := time.Duration(n) * time.Millisecond

	fmt.Println(timeout)

	select {
	case <-time.After(timeout):
		c <- "remote data: " + successText
	case <-ctx.Done():
		fmt.Println("dont send ", successText)
		stopChan <- true
	}
}

