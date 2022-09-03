package concurrency

import "fmt"

func blockingChannel(myChannel chan string) {

	myChannel <- "foobar"

	x := <-myChannel
	fmt.Println(x)
}
