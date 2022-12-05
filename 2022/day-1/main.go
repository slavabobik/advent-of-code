package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var max, sum int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d", max)
}
