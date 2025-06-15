package headless

import (
	"flag"
	"fmt"
	"job-visualizer/pkg/gui"
	"job-visualizer/pkg/shared"
	"strings"
)

func CheckCLIArguments() bool {
	headless := flag.Bool("headless", false, "Run in headless mode (no GUI)")
	flag.Parse()
	return *headless
}

func RunGUIorHeadless(headless bool, allJobData []shared.JobData) {
	if headless {
		for i, job := range allJobData {
			if i%100 == 0 {
				fmt.Printf("%-4s | %-25s | %-55s | %-25s\n",
					"#", "Location", "Job Title", "Company Name")
				fmt.Println(strings.Repeat("-", 120))
			}
			fmt.Printf("%-4d | %-25s | %-55s | %-25s\n",
				i+1, job.Location, job.JobTitle, job.CompanyName)
		}
	} else {
		gui.CreateGui(allJobData)
	}
}
