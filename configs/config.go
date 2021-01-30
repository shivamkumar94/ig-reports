package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Configs struct {
	SpreadsheetID         string `json:"SpreadsheetID"`
	SheetNameWithRange    string `json:"SheetNameWithRange"`
	ResearchJRSheetName   string `json:"ResearchJRSheetName"`
	SessionIDSheetName    string `json:"SessionIDSheetName"`
	SessionId             string `json:"SessionId"`
	IGRDatabaseSheetName  string `json:"IGRDatabaseSheetName"`
	NOSSearchSheetName    string `json:"NOSSearchSheetName"`
	NOSSearch2SheetName   string `json:"NOSSearch2SheetName"`
	NOSSearch3SheetName   string `json:"NOSSearch3SheetName"`
	NOSSearch4SheetName   string `json:"NOSSearch4SheetName"`
	NOSDashboardSheetName string `json:"NOSDashboardSheetName"`
}

var (
	Configurations = Configs{}
)

func SetConfig() {
	input, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	error := json.Unmarshal(input, &Configurations)
	if error != nil {
		fmt.Println("Config file is missing in root directory")
		panic(error)
	} else {
		fmt.Println("Follwing values has been picked from config values:")
		fmt.Println(Configurations)
	}
}
