// define maps
package main

import "fmt"

var str_key string = "unknown"   //for check key in a map
var str_value string = "unknown" //for print value in a map
var exist_date bool = false      //for print key status in a map

func main() {
	/* first style define a map
	elements := map[string]string{ //[key]value
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	*/
	elements := make(map[string]string) //second style define a map
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Baron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	str_key = "F"

	if e, ok := elements[str_key]; ok { //return 2 parameters: value of map and bool if request - good
		str_value = e
		exist_date = ok
	}
	fmt.Println("Key:[", str_key, "] value:", str_value, ";date exist:", exist_date)
}
