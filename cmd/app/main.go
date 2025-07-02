package main

import (
	"job-visualizer/pkg/gui"
	"job-visualizer/pkg/headless"
)

func main() {
	isHeadless := headless.CheckCLIArguments()
	gui.RunGUIorHeadless(isHeadless)
}
