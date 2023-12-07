package main

/*

--- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky") and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky ("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

To begin, get your puzzle input.

*/

import (
	"fmt"
	"strconv"
	"strings"
)

// func main() {
// 	calibrationDocumentSummation()
// }

func calibrationDocumentSummation() int {
	var result int

	

	// ledger := []string {	
	// 	"lczpbbjnh1fourxbskj",
	// }

	ledger := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	for _, str := range ledger {
		event := parse(str)
		// fmt.Println("parse: ", event)
		result += event
	}
	fmt.Println(result)
	return result
}


func parse(target string) int {	
	dictionary := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
	
	length := len(target)
	
	onesCache := -1
	tensCache := -1

	foreword := ""
	backword := ""
	for i, j := 0, length - 1; i <= j; {
		if tensCache == -1 {
			if target[i] > 57 {
				if found, ok := dictionary[foreword]; !ok {
					foreword += string(target[i])
				} else {
					tensCache = found
				}
			}

			for k, v := range dictionary {
				if strings.Contains(foreword, k) {
					// fmt.Println("inside guard with forword: ", foreword)
					tensCache = v
					fmt.Println(tensCache)
				}
			}

			if target[i] <= 57 {
				first, err := strconv.Atoi(string(target[i]))
				if err != nil {
					panic(err)
				}

				tensCache = first
			}

			i++
		}

		if onesCache == -1 {
			if target[j] > 57 {
				if found, ok := dictionary[backword]; !ok {
					container := string(target[j])
					container += backword
					backword = container
				} else {
					onesCache = found
				}
			}

			for k, v := range dictionary {
				if strings.Contains(backword, k) {
					fmt.Println("inside guard with backword: ", backword)
					onesCache = v
					fmt.Println(onesCache)
				}
			}

			if target[j] <= 57 {
				first, err := strconv.Atoi(string(target[j]))
				if err != nil {
					panic(err)
				}

				onesCache = first
			}

			j--
		}


		if tensCache != -1 && onesCache != -1 {
			break
		}
	}

	if onesCache == -1 && tensCache != -1 {
		// fmt.Println("single cache tens: ", (tensCache * 10) + tensCache, tensCache, onesCache, " ", target)
		return (tensCache * 10) + tensCache
	}

	if onesCache != -1 && tensCache == -1 {
		// fmt.Println("single cache ones: ", (onesCache * 10) + onesCache, onesCache, tensCache, " ", target)
		return (onesCache * 10) + onesCache
	}

	// fmt.Println("full cache: ", (tensCache * 10) + onesCache, tensCache, onesCache, " ",  target)
	return (tensCache * 10) + onesCache
}
