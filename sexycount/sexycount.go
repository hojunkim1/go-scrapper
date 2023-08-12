package sexycount

import (
	"fmt"
	"time"
)

// CheckIsSexy check is sexy
func CheckIsSexy() {
	people := [2]string{"nico", "flynn"}

	c := make(chan string)

	for _, person := range people {
		go isSexy(person, c)
	}

	for range people {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + " is sexy"
}
