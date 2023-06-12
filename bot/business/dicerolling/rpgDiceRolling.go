package dicerolling

import (
	"sort"
	"strconv"
	"strings"
)

// InitiativeRoller
//Input:
//Give it the amount of individuals you are rolling for,
//AND their respective modifier
//
//Output:
//Returns their initiative roll in order, with their name if needed.
func InitiativeRoller(input string) string {
	//Map created
	nameToTotalRollMap := make(map[string]int)
	nameToModMap := make(map[string]string)
	nameToFlatRollMap := make(map[string]int)
	//Input sliced into an array
	slicedInput := strings.Split(input[12:], ",")

	//Loop through the array,
	//Then, populate map with rolled value and name pairing.
	for i, value := range slicedInput {
		value = strings.TrimSpace(value)
		if i%2 != 0 {
			nameToModMap[slicedInput[i-1]] = value
			numericalValue, err := strconv.Atoi(value)
			if err != nil {
				return "Error decoding modifier value"
			}
			rollValue := randomNumberGenerator(20)
			nameToFlatRollMap[slicedInput[i-1]] = rollValue
			nameToTotalRollMap[slicedInput[i-1]] = rollValue + numericalValue
		}
	}

	// Create slice of key-value pairs
	pairs := make([][2]interface{}, 0, len(nameToTotalRollMap))

	for k, v := range nameToTotalRollMap {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort slice based on values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1].(int) > pairs[j][1].(int)
	})

	// Extract sorted keys
	keys := make([]string, len(pairs))
	for i, p := range pairs {
		keys[i] = p[0].(string)
	}

	//Append onto a message
	outputString := "```ansi\n		Initiative:\n"
	shortOrderString := "\n\n		IN ORDER: \n"
	for i, k := range keys {
		outputString += strings.TrimSpace(k) + "	Roll: " + strconv.Itoa(nameToFlatRollMap[k]) + " 		Modifier:" + nameToModMap[k] + "		Total Roll:		[\u001B[37m" + strconv.Itoa(nameToTotalRollMap[k]) + "\u001B[0m]" + "\n"
		shortOrderString += strconv.Itoa(i+1) + "	" + strings.TrimSpace(k) + "\n"
	}
	outputString += shortOrderString
	outputString += "\n```"

	return outputString

}
