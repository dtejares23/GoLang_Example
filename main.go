package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
)

type Todo struct {
	Zone     string `json:"AD"`
	Ae       string `json:"AD"`
	UsCarDuk string `json:"US-CAR-DUK"`
}

// func api(url string) {
// 	response, err := http.Get(url)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(response)
// }

func main() {

	// url := os.Args[1]
	// api(url)

	// api =: api.GetZones()
	// fmt.Println(api)

	//START OF CALLING API
	fmt.Println("1. Performing Http Get...")
	resp, err := http.Get("https://api.electricitymap.org/v3/zones")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	// Convert response body to Todo struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
	//END OF CALLING API

	fmt.Println("START OF ANOTHER STRING")
	fmt.Println("Hello, Daryl...")

	var e int
	e = 5
	fmt.Println(e)

	var f float32 = 5.4
	fmt.Println(f)

	firstName := "Daryl"
	fmt.Println(firstName)

	lastName := "Tejares"
	fmt.Println(lastName)

	ellie := true
	fmt.Println(ellie)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	aliasName := "Ellie"
	fmt.Println(aliasName)

	pointerName := &aliasName
	fmt.Println(pointerName, *pointerName)

	aliasName = "Ellie2"
	fmt.Println(pointerName, *pointerName)
}
