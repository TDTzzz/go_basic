package goroutine

import "log"

func main() {

	for i := 1; i <= 100; i++ {
		go output(i)
	}
}

func output(num int) {
	log.Println(num)
}
