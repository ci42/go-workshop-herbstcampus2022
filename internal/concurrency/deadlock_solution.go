package concurrency

import "fmt"

func blockingChannel(myChannel chan string) {

	go func() {
		myChannel <- "foobar"
	}()

	x := <-myChannel
	fmt.Println(x)
}
