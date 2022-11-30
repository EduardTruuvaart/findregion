package main

import (
	"embed"
	"encoding/json"
	"findregion/dto"
	"fmt"
)

//go:embed ip-ranges.json
var content embed.FS

func main() {
	textBytes, _ := content.ReadFile("ip-ranges.json")

	var result dto.IPRanges
	err := json.Unmarshal(textBytes, &result)

	if err != nil {
		panic(err)
	}

	for _, prefix := range result.Prefixes {
		fmt.Println(prefix)
	}

	// jsonFile, err := os.Open(filePath)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully Opened users.json")
	// // defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// fmt.Printf("byteValue len %d\n", len(byteValue))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(byteValue), &result)

	// fmt.Println(result["prefixes"])

}
