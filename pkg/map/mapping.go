package mapping

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"io"
	"job-visualizer/pkg/shared"
	"net/http"
	"net/url"
	"strconv"
)

func ProcessLatLongs(jobs []shared.JobData) []shared.JobData {
	shared.CachedLocations = make(map[string]shared.LatLong)
	jobs = standardizeLocations(jobs)
	getLatLongs(jobs)
	jobs = assignLatLongs(jobs)
	return jobs
}

func standardizeLocations(jobs []shared.JobData) []shared.JobData {
	reNumbers := regexp.MustCompile(`[0-9]+`)
	rePunctuation := regexp.MustCompile(`[^\w\s]`)
	for i, job := range jobs {
		location := job.Location
		location = strings.ToLower(location)
		location = reNumbers.ReplaceAllString(location, "")
		location = rePunctuation.ReplaceAllString(location, "")
		location = strings.Join(strings.Fields(location), " ")
		jobs[i].Location = location
	}
	return jobs
}

func getLatLongs(jobs []shared.JobData) {
	for i, job := range jobs {
		fmt.Println(job.Location)
		fmt.Printf("\rcaching job locations (%d/%d)\n", i+1, len(jobs))
		if _, ok := shared.CachedLocations[job.Location]; ok {
		} else {
			responseBody := getNominatimResponse(job.Location)

			var locations []shared.JsonLocation
			err := json.Unmarshal(responseBody, &locations)
			shared.CheckErrorWarn(err)
			if len(locations) > 0 {
				addLocationToCache(job.Location, locations)

			}
		}
	}
}

func assignLatLongs(jobs []shared.JobData) []shared.JobData {
	for i, job := range jobs {
		if coordinates, ok := shared.CachedLocations[job.Location]; ok {
			jobs[i].LatLong = coordinates
		}
	}
	return jobs
}

func getNominatimResponse(location string) []byte {
	encodedLocation := url.QueryEscape(location)
	apiUrl := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", encodedLocation)
	response, err := http.Get(apiUrl)
	shared.CheckErrorWarn(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		shared.CheckErrorWarn(err)
	}(response.Body)
	responseBody, err := io.ReadAll(response.Body)
	shared.CheckErrorWarn(err)
	return responseBody
}

func addLocationToCache(jobLocation string, locations []shared.JsonLocation) {
	latitude, err := strconv.ParseFloat(locations[0].Lat, 64)
	shared.CheckErrorWarn(err)
	longitude, err := strconv.ParseFloat(locations[0].Lon, 64)
	shared.CheckErrorWarn(err)
	shared.CachedLocations[jobLocation] = shared.LatLong{
		Latitude:  latitude,
		Longitude: longitude,
	}
}
