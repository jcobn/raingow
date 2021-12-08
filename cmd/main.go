package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/jcobn/raingow/utils"
)

func main() {
	flag.Parse()
	if *flClearTerminal {
		utils.ClearTerminal()
	}
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		utils.SendUsageMessage()
		return
	}
	bytes, _ := ioutil.ReadAll(os.Stdin)
	input := string(bytes)
	utils.PrintRgbText(input)
}

var (
	flClearTerminal = flag.Bool("c", false, "Clears the terminal when printing text")
)
