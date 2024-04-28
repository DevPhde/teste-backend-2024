package main

import (
	"ms-go/app/kafka/consumers"
	"ms-go/router"
)

func main() {

	go func() {
		consumers.ProductConsumer()
	}()
	router.Run()

}
