package geo

import (
    "net/http"
	"io/ioutil"
	"fmt"
	"github.com/yalp/jsonpath"
	"encoding/json"
	"os"
)

// type location struct {
// 	lat string
// 	lon string
// }

func GetLocation(country string, storeId string) interface{} {

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
	err1 := json.Unmarshal(bodyBytes, &store)
	location, err1 := jsonpath.Read(store, "$..location")
	
	if err1 != nil {
        panic(err1)
    }

	fmt.Println(location)

	return location
}