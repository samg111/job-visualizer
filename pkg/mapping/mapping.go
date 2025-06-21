package mapping

import (
	"fmt"
	"job-visualizer/pkg/shared"
	"net/http"

	"github.com/morikuni/go-geoplot"
	"github.com/skratchdot/open-golang/open"
)

var geoplotMap *geoplot.Map

func init() {
	http.HandleFunc("/map", mapPage)
	http.HandleFunc("/innerMap", innerMap)
	// http.Handle("/map", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	// 	if geoplotMap != nil {
	// 		err := geoplot.ServeMap(writer, request, geoplotMap)
	// 		shared.CheckErrorWarn(err)
	// 	} else {
	// 		http.Error(writer, "Map not ready", http.StatusServiceUnavailable)
	// 	}
	// }))
}

func mapPage(writer http.ResponseWriter, request *http.Request) {
	if geoplotMap != nil {
		writer.Header().Set("Content-Type", "text/html")
		fmt.Fprint(writer, `
            <html>
                <head>
                    <title>job-visualizer Map</title>
                </head>
                <body style="margin:0;padding:0;">
                    <iframe src="/innerMap" style="width:100vw;height:100vh;border:none;"></iframe>
                </body>
            </html>
        `)
	} else {
		http.Error(writer, "Map not ready", http.StatusServiceUnavailable)
	}
}

func innerMap(writer http.ResponseWriter, request *http.Request) {
	if geoplotMap != nil {
		err := geoplot.ServeMap(writer, request, geoplotMap)
		shared.CheckErrorWarn(err)
	} else {
		http.Error(writer, "Map not ready", http.StatusServiceUnavailable)
	}
}

func GenerateMap(jobs []shared.JobData) []shared.JobData {
	geoplotMap = createGeoplotMap(jobs)
	shared.Window.Server = createHttpServer()
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

func createHttpServer() *http.Server {
	if shared.Window.Server != nil {
		err := shared.Window.Server.Close()
		shared.CheckErrorWarn(err)
	}
	server := &http.Server{Addr: ":8080"}
	go func() {
		err := server.ListenAndServe()
		shared.CheckErrorWarn(err)
	}()
	return server
}

func openWebpage() {
	url := "http://localhost:8080/map"
	err := open.Run(url)
	shared.CheckErrorWarn(err)
}
