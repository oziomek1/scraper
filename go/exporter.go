package main

import (
	"os"
	"log"
	"encoding/csv"
	"reflect"
)

func exportData(offers []Offer, filename string) {
	file, err := os.Create(filename + ".csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	data := convertType(offers)

	writeFile(writer, data)
}

func parse(data Params) []string {
	dataValue := reflect.ValueOf(data)
	n := dataValue.NumField()

	content := make([]interface{}, n)
	for i := 0; i < n; i++ {
		x := dataValue.Field(i)
		if x.CanInterface() {
			s := x.Interface()
			content[i] = s
		}
 		content[i] = x.String()
	}
	var array []string
	for _, value := range content {
		array = append(array, value.(string))
	}
	return array
}

func convertType(offers []Offer) ([][]string) {
	var data [][]string
	for _, offer := range offers {
		line := parse(offer.params)
		data = append(data, line)
	}
	return data
}

func writeFile(writer *csv.Writer, data [][]string)  {
	writeHeader(writer)

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func writeHeader(writer *csv.Writer)  {
	// Has to do it manually as golang map order is randomized
	headers := []string{"URL", "ID", "PRICE", "CURRENCY",
						"SELLER", "CATEGORY", "MAKE", "MODEL",
						"GENERATION", "YEAR", "MILEAGE", "ENGINE_CAPACITY",
						"FUEL_TYPE", "POWER", "GEARBOX", "POWERTRAIN",
						"CHASSIS", "DOORS", "SEATS",
						"COLOUR", "ACRYLIC", "METALLIC", "PEARL",
						"LICENCE_PLATE", "VIN", "CRASHED", "VAT_INVOICE",
						"FINANCING", "FIRST_OWNER", "COLLISION_FREE",
						"RHD", "COUNTRY",
						"FIRST_REGISTRATION", "REGISTER_IN_POLAND", "SERVICE_AUTHORISED", "ANTIQUE_REGISTERED",
						"TUNING", "CONDITION"}
	err := writer.Write(headers)
	checkError("Cannot add header to file", err)
}

func checkError(message string, err error)  {
	if err != nil {
		log.Fatal(message, err)
	}
}