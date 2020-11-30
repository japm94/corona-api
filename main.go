package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Data struct used to encode csv data to json
type Data struct {
	CountryRegion string `json:"Country_Region"`
	LastUpdate    string `json:"Last_Update"`
	Lat           string `json:"Lat"`
	Long          string `json:"Long_"`
	Confirmed     string `json:"Confirmed"`
	Deaths        string `json:"Deaths"`
}

func main() {

	jasondata := convertCSVToJSON()
	fmt.Println(string(jasondata))

}

func todaysDate() string {
	today := time.Now().Format("01-02-2006")

	return today
}

// Convert CSV to Data
func convertCSVToJSON() []byte {

	var oneRecord Data
	var allRecords []Data
	records := ReadCSV("https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_daily_reports/" + todaysDate() + ".csv")
	for _, rec := range records {
		oneRecord.CountryRegion = rec[3]
		oneRecord.Lat = rec[5]
		oneRecord.Long = rec[6]
		oneRecord.Confirmed = rec[7]
		oneRecord.Deaths = rec[8]
		allRecords = append(allRecords, oneRecord)
	}

	jasondata, err := json.Marshal(allRecords)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jasondata))

	return jasondata
}

// ReadCSV read the csv and return slice of slices with all data
func ReadCSV(url string) [][]string {

	csvFile, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Body.Close()

	reader := csv.NewReader(csvFile.Body)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	if len(records) < 1 {
		log.Fatal("Something wrong, the file maybe empty or length of the lines are not the same")
	}
	headersArr := make([]string, 0)
	for _, headE := range records[0] {
		headersArr = append(headersArr, headE)
	}
	// Removing header
	records = records[1:]
	fmt.Println(records)
	return records
}
