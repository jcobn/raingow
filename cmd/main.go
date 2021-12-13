package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/jcobn/raingow/utils"
)

func main() {
	flag.Parse()
	info, _ := os.Stdin.Stat()
	input := ""

	if info.Mode()&os.ModeCharDevice != 0 {
		if len(*flInput) > 0 {
			input = *flInput
		} else {
			utils.SendUsageMessage()
			return
		}
	} else {
		bytes, _ := ioutil.ReadAll(os.Stdin)
		input += string(bytes)
	}
	if *flClearTerminal {
		utils.ClearTerminal()
	}
	utils.PrintRgbText(input)
}

var (
	flClearTerminal = flag.Bool("c", false, "Clears the terminal when printing text")
	flInput         = flag.String("i", "", "Input string")
)
