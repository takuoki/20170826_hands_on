package main

// goroutine test

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func waitRandom(max int) int {
	i := rand.Intn(max)
	time.Sleep(time.Duration(i) * time.Second)
	return i
}

func exec(l string) {
	i := waitRandom(5)
	log.Printf("exec   : %s (wait time = %d sec)\n", l, i)
}

func execForGoroutine(ch chan string, l string) {
	exec(l)
	ch <- "finish : " + l
}

func main() {

	goroutine := true

	log.Print("*** start ***")

	// read file
	var fp *os.File
	fp, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	ch := make(chan string)
	count := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		l := scanner.Text()
		log.Printf("call   : %s\n", l)

		if goroutine {
			go execForGoroutine(ch, l)
			count++
		} else {
			exec(scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if goroutine {
		log.Print("*** finish call ***")

		// wait to finish
		for i := 0; i < count; i++ {
			log.Println(<-ch)
		}
	}

	log.Print("*** finish all ***")
}
