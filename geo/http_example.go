package geo

import (
    "net/http"
	"io/ioutil"
	"encoding/json"
	"os"
)
type Response struct {
	Location Location
}

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

	var store = Response{}
	json.Unmarshal([]byte(bodyBytes), &store)

	return store.Location
}