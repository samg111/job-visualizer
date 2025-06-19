package mapping

import (
	"fmt"
	"job-visualizer/pkg/jobdata/processing"
	"job-visualizer/pkg/shared"
	"net/http"

	"github.com/morikuni/go-geoplot"
	"github.com/skratchdot/open-golang/open"
)

// var serverCount int

func GenerateMap(jobs []shared.JobData) []shared.JobData {
	jobs = processing.ProcessLatLongs(jobs)
	geoplotMap := createGeoplotMap(jobs)
	shared.Window.Server = createHttpServer(geoplotMap)
	openWebpage()
	return jobs
}

func createGeoplotMap(jobs []shared.JobData) *geoplot.Map {
	boston := &geoplot.LatLng{
		Latitude:  42.361145,
		Longitude: -71.057083,
	}
	geoplotMap := &geoplot.Map{
		Center: boston,
		Zoom:   7,
		Area: &geoplot.Area{
			From: boston.Offset(-0.1, -0.1),
			To:   boston.Offset(0.2, 0.2),
		},
	}
	commonLocations := make(map[shared.LatLong]string)
	for _, job := range jobs {
		if _, ok := commonLocations[job.LatLong]; ok {
			commonLocations[job.LatLong] = commonLocations[job.LatLong] + fmt.Sprintf("%s\n", job.CompanyName)
		} else {
			commonLocations[job.LatLong] = fmt.Sprintf("%s\n", job.CompanyName)
		}
	}
	for key, value := range commonLocations {
		latitude := key.Latitude
		longitude := key.Longitude
		coordinates := &geoplot.LatLng{
			Latitude:  latitude,
			Longitude: longitude,
		}
		icon := geoplot.ColorIcon(255, 255, 0)
		geoplotMap.AddMarker(&geoplot.Marker{
			LatLng:  coordinates,
			Popup:   value, //clicked description
			Tooltip: value, //hover word
			Icon:    icon,
		})
	}
	return geoplotMap
}

func createHttpServer(geoplotMap *geoplot.Map) *http.Server {
	if shared.Window.Server != nil {
		err := shared.Window.Server.Close()
		shared.CheckErrorWarn(err)
	}
	// serverCount++
	server := &http.Server{Addr: ":8080"}
	// pattern := fmt.Sprintf("/%d", serverCount)
	http.Handle("/map", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := geoplot.ServeMap(writer, request, geoplotMap)
		shared.CheckErrorWarn(err)
	}))
	go func() {
		err := server.ListenAndServe()
		shared.CheckErrorWarn(err)
	}()
	return server
}

func openWebpage() {
	// url := fmt.Sprintf("http://localhost:8080/%d", serverCount)
	url := "http://localhost:8080/map"
	err := open.Run(url)
	shared.CheckErrorWarn(err)
}
