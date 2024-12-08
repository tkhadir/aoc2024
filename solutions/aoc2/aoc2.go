package main

import (
    "log"
	"aoc.com/utils"
	"strings"
	"strconv"
)

func checkElementsOrder(report string, increasing bool, idxIgnore int) bool {
	previous := 0
	elements := strings.Split(report, " ")
	for index2, ele2 := range elements {
		value, _ := strconv.Atoi(ele2)
		if (index2 > 0) &&
			index2 != idxIgnore {
			if (increasing && previous < value) || 
					((!increasing) && previous > value) {
				//log.Println("compare ", previous, " and ", value, " true", " with index ", index2)
				previous = value
			} else {
				//log.Println("compare ", previous, " and ", value, " false", " with index ", index2)
				return false
			}
		} else if (index2 == 0) {
			previous = value
		}
	}
	return true
}

func checkElementsOrderDiff(report string, idxIgnore int) bool {
	previous := 0
	elements := strings.Split(report, " ")
	diff_ := [...]int{1, 2, 3}
	for index2, ele2 := range elements {
		value, _ := strconv.Atoi(ele2)
		if (index2 > 0) &&
			index2 != idxIgnore {
			vdiff := 0
			if (previous > value) {
				vdiff = previous - value
			} else {
				vdiff = value - previous
			}
			if exists_(diff_, vdiff) {
				previous = value
			} else {
				return false
			}
		} else if (index2 == 0) {
			previous = value
		}
	}
	return true
}

func exists_(arr_ [3]int, find_ int) bool {
	for _, v := range arr_  {
		if (v == find_) {
			return true
		}
	}
	return false
}

func step1() {
	log.Println("start solving")
	inputs := utils.ReadFile("../../inputs/aoc2.txt")
	data := strings.Split(inputs, "\n")

	total := 0

	for _, ele := range data {
		//report := strings.Split(ele, " ")
		//log.Println("analyzing report ", ele)

		isInc := checkElementsOrder(ele, true, -1)
		//log.Println(report, " is inc ", isInc)
		isDec := checkElementsOrder(ele, false, -1)
		//log.Println(report, " is dec ", isDec)
		isOrderedDiff := checkElementsOrderDiff(ele, -1)

		isOrdered := isInc || isDec
		//log.Println(report, " is ordered ", isOrdered)
		//log.Println(report, " is ordered diff", isOrderedDiff)

		isSafe := isOrdered && isOrderedDiff
		//log.Println(report, " is safe ", isSafe)

		if (isSafe) {
			total += 1
		}
	}

	log.Println("end solving ", total)
}

func step2() {
	log.Println("start solving")
	inputs := utils.ReadFile("../../inputs/aoc2.txt")
	data := strings.Split(inputs, "\n")

	total := 0

	for _, ele := range data {
		//report := strings.Split(ele, " ")
		//log.Println("analyzing report ", ele)

		isInc := checkElementsOrder(ele, true, -1)
		//log.Println(report, " is inc ", isInc)
		isDec := checkElementsOrder(ele, false, -1)
		//log.Println(report, " is dec ", isDec)
		isOrderedDiff := checkElementsOrderDiff(ele, -1)

		isOrdered := isInc || isDec
		//log.Println(report, " is ordered ", isOrdered)
		//log.Println(report, " is ordered diff", isOrderedDiff)

		isSafe := isOrdered && isOrderedDiff
		//log.Println(report, " is safe ", isSafe)

		if (isSafe) {
			total += 1
		} else {
			for idx, _ := range strings.Split(ele, " ") {
				//log.Println("check idx ", idx)
				//report := strings.Split(ele, " ")
				//log.Println("analyzing report ", ele)

				isInc := checkElementsOrder(ele, true, idx)
				//log.Println(report, " is inc ", isInc)
				isDec := checkElementsOrder(ele, false, idx)
				//log.Println(report, " is dec ", isDec)
				isOrderedDiff := checkElementsOrderDiff(ele, idx)

				isOrdered := isInc || isDec
				//log.Println(report, " is ordered ", isOrdered)
				//log.Println(report, " is ordered diff", isOrderedDiff)

				isSafe := isOrdered && isOrderedDiff
				//log.Println(report, " is safe ", isSafe)

				if (isSafe) {
					//log.Println("good idx ", idx)
					total += 1
					break
				}
			}
		}
		//log.Println("is ", total)
	}

	log.Println("end solving ", total)
}

func main() {
    log.Println("solve step1")
	step1()

	log.Println("solve step2")
	step2()
}