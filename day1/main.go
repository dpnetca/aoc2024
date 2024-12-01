package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftList := []int{}
	rightList := []int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "   ")
		leftNumber, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Println(err)
		}
		rightNumber, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Println(err)
		}
		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)
	difTotal := 0
	simTotal := 0
	for i, v := range leftList {
		difTotal += int(math.Abs(float64(v - rightList[i])))
		count := 0
		for _, rv := range rightList {
			if v == rv {
				count++
			}
		}
		simTotal += v * count

	}
	fmt.Printf("difference: %v\n", difTotal)
	fmt.Printf("simularity: %v\n", simTotal)

}
