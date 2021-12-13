package utils

func SendUsageMessage() {
	PrintRgbText("The command is intended to work with pipes or the -i (input) flag.\nUsage: <stdin> | raingow (-c)\n  OR: raingow -i <input string> (-c)")
}
