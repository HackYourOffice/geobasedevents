package main

import g "../geo"
import  "fmt"

func main() {

	var location = g.GetLocation("de", "432")
	fmt.Println(location)

}