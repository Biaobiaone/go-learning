package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var answer int = rand.Intn(100) + 1
	var guess int = 50
	var min, max = 1, 100
	var count = 0
	for guess != answer {
		if guess > answer {
			fmt.Printf("%v不是正确答案，猜大了\n", guess)
			max = guess
			guess = (max + min) / 2
		} else {
			fmt.Printf("%v不是正确答案，猜小了\n", guess)
			min = guess
			guess = (max + min) / 2
		}
		count++
	}
	count++
	fmt.Printf("猜对了，正确答案是%v,一共猜了%v次", answer, count)
}
