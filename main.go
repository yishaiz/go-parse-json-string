package main

import (
	"encoding/json"
	"fmt"
)

const (
	payload = `{"jwt":"fdnvjkdsnjkvmsjkdlmvsjkdfmjnv", "expires_at": "2021-05-05T07:20:11Z" }`
)

func main() {
	fmt.Println("xxxxxx")
	fmt.Println(payload)

	birdJson := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var birdsResult map[string]interface{}
	var jsonResult map[string]string
	err1 := json.Unmarshal([]byte(payload), &jsonResult)
	if err1 != nil {
		fmt.Println("Error in parsing payload")
	}

	fmt.Println(jsonResult["jwt"])

	err2 := json.Unmarshal([]byte(birdJson), &birdsResult)
	if err2 != nil {
		fmt.Println("Error in parsing birds json")
	}
	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := birdsResult["birds"].(map[string]interface{})

	for key, value := range birds {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}

}
