package main

import (
	"fmt"
)

func main() {
	loop()
}

func loop() {
	for x := 1; x <= 5; x++ {
		for i := 1; i <= 5; i++ {
			if i == 1 || i == 5 {
				fmt.Print("0")
			} else if x == 1 || x == 5 {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}

}
