package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"wan-api-kol-event/Initializers"
	"wan-api-kol-event/Models"
)

func init() {
	Initializers.LoadEnvironmentVariables()
	Initializers.ConnectToDB()
}

func main() {
	kols := getJsonData("Migrate/dummiesData.json")
	fmt.Println(kols)
}

func getJsonData(filePath string) []Models.Kol {
	var kols []Models.Kol
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Panicln("Fail to open json file", err)
	}
	jsonErr := json.Unmarshal(fileBytes, &kols)
	if jsonErr != nil {
		log.Panicln("Fail to parse json to KOL", jsonErr)
	}
	return kols
}