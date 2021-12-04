package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	var people = [5]string{"zero", "hero", "kero", "bero", "dero"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is sexy"
}
