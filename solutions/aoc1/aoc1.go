package main

import (
    "log"
	"aoc.com/utils"
	"strings"
	"strconv"
)

func delete_at_index(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func minFind(arr []int) int {
	min := arr[0]
	minIndex := 0
	for point, num1 := range arr {
	   if num1 < min {
		  min = num1
		  minIndex = point
	   }
	}
	return minIndex
}

func countFind(arr []int, ele int) int {
	count := 0
	for _, num1 := range arr {
	   if ele == num1 {
		  count += 1
	   }
	}
	return count
}

func step1() {
	total := 0
	var left []int
	var right []int

	inputs := utils.ReadFile("../../inputs/aoc1.txt")
	data := strings.Split(inputs, "\n")

	for index, ele := range data {
		log.Println(index, " - ", ele)
		locations := strings.Split(ele, " ")

		log.Println("loc1 ", locations[0])
		loc1, err1 := strconv.Atoi(locations[0])
		if err1 != nil {
			log.Println("err1 conv ", loc1)
		}

		log.Println("loc2 ", locations[3])
		loc2, err2 := strconv.Atoi(locations[3])
		if err2 != nil {
			log.Println("err2 conv ", loc2)
		}

		left = append(left, loc1)
		right = append(right, loc2)
	}

	for len(left) > 0 {
		
		rang1 := minFind(left)
		rang2 := minFind(right)

		d := 0
		if left[rang1] > right[rang2] {
				d = left[rang1] - right[rang2]
		} else {
				d = right[rang2] - left[rang1]
		}

		log.Println("distance ", d)
		total += d

		left = delete_at_index(left, rang1)
		right = delete_at_index(right, rang2)
	}

	log.Println("total == ", total)
}

func step2() {
	total := 0
	var left []int
	var right []int

	inputs := utils.ReadFile("../../inputs/aoc1.txt")
	data := strings.Split(inputs, "\n")

	for index, ele := range data {
		log.Println(index, " - ", ele)
		locations := strings.Split(ele, " ")

		log.Println("loc1 ", locations[0])
		loc1, err1 := strconv.Atoi(locations[0])
		if err1 != nil {
			log.Println("err1 conv ", loc1)
		}

		log.Println("loc2 ", locations[3])
		loc2, err2 := strconv.Atoi(locations[3])
		if err2 != nil {
			log.Println("err2 conv ", loc2)
		}

		left = append(left, loc1)
		right = append(right, loc2)
	}

	for _, ele := range left {
		occur := countFind(right, ele)

		total += (ele * occur)
	}

	log.Println("total == ", total)
}

func main() {
	log.Println("start aoc1")

	step1()

	step2()
}