// map_test1.go
package main

import (
	"fmt"
)

func main() {
	// var countryCapitalMap map[string]string //gen map,无序
	countryCapitalMap := make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["United States"] = "New York"

	for country := range countryCapitalMap {
		fmt.Printf("The capital of %s is %s\n", country, countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["United States"] //查看元素在集合中是否存在
	if ok {
		fmt.Println("Capital of United States is ", capital, "!!!")
	} else {
		fmt.Println("Capital of United States is not present.")
	}

	delete(countryCapitalMap, "United States")
	fmt.Println("United States has been deleted.")
}
