package dicerolling

import (
	"math/rand"
	"sort"
	"strconv"
)

func FiveEStats() (string, error) {
	message := "```ansi\n"
	rolledList := [6][5]int{}
	statHolder := [5]int{}
	messageHolder := [6]string{}
	totalHolder := [6]int{}
	for i := 0; i <= 5; i++ {
		statHolder[0] = rand.Intn(6) + 1
		statHolder[1] = rand.Intn(6) + 1
		statHolder[2] = rand.Intn(6) + 1
		statHolder[3] = rand.Intn(6) + 1
		sort.Ints(statHolder[:])

		messageHolder[i] += "Rolls: "
		for j := 0; j < 4; j++ {
			if statHolder[j] == 6 {
				messageHolder[i] += "[\u001B[32m" + strconv.Itoa(statHolder[j]) + "\u001B[0m] "
			} else if statHolder[j] == 1 {
				messageHolder[i] += "[\u001B[31m" + strconv.Itoa(statHolder[j]) + "\u001B[0m] "
			} else {
				messageHolder[i] += "[" + strconv.Itoa(statHolder[j]) + "] "
			}
			if j != 0 {
				statHolder[4] += statHolder[j]
			}
			if j == 3 {
				messageHolder[i] += "    Total: " + strconv.Itoa(statHolder[4]) + "\n"
				totalHolder[i] += statHolder[4]
			}
		}
		rolledList[i] = statHolder
		statHolder[4] = 0
	}
	sort.Ints(totalHolder[:])
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			x, _ := strconv.Atoi(messageHolder[j][len(messageHolder[j])-3 : len(messageHolder[j])-1])
			if totalHolder[5-i] == x {
				message += messageHolder[j]
			}
		}
	}
	message += "```"
	return message, nil
}
