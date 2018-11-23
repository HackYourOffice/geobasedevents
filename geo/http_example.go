package geo

import (
    "net/http"
	"io/ioutil"

	"github.com/yalp/jsonpath"
	"encoding/json"
	"os"

)

 type Location struct {
 	Lat float64
	Lon float64
}

func GetLocation(country string, storeId string) Location {

	geoBasedEventUrl := os.Getenv("GEO_BASED_EVENT_URL")

    rs, err := http.Get(geoBasedEventUrl + country + "/" + storeId)
    // Process response
    if err != nil {
        panic(err) // More idiomatic way would be to print the error and die unless it's a serious error
    }	
    defer rs.Body.Close()

    bodyBytes, err := ioutil.ReadAll(rs.Body)
    if err != nil {
        panic(err)
    }

	//bodyString := string(bodyBytes)
	var store interface{}
	err1 := json.Unmarshal([]byte(bodyBytes), &store)
	//location :=  new(Location)
	rawLat, err1 := jsonpath.Read(store, "$.location.lat")
	rawLon, err1 := jsonpath.Read(store, "$.location.lon")

	if err1 != nil {
        panic(err1)
	}
	
	lat := rawLat.(float64)
	lon := rawLon.(float64)

	var location Location = Location{lat, lon}

	return location
}