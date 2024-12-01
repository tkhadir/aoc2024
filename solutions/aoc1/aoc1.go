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

func main() {
	log.Println("start aoc1")

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

	log.Println("total == ", total)
}