package commons

func ProbInstruction() string {
	message := "```ansi\n"
	message += "\nThe format goes as: \n\n!p modifier,difficultyClass\n"
	message += "or \n!p mod,dc\n"
	message += "```"
	return message
}
