package mapping

import (
	"job-visualizer/pkg/mapping/processing"
	"job-visualizer/pkg/shared"
)

func GenerateMap(jobs []shared.JobData) []shared.JobData {
	jobs = processing.ProcessLatLongs(jobs)
	return jobs
}
