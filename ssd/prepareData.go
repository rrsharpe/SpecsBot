package ssd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/go-zoox/fetch"
)

func readCsvFile() [][]string {
	response, err := fetch.Get("https://docs.google.com/spreadsheets/d/1B27_j9NDPU3cNlj2HKcrfpJKHkOf-Oi1DbuuQva2gT4/export?format=csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(bytes.NewReader(response.Body))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for ", err)
	}

	return records
}

var bannedWords mapset.Set = mapset.NewSet( // SPECIFICALLY for titles/models
	"",       // Some product page cells are empty
	"BLACK",  // WD Black models don't always have Black listed in the table
	"GAMING", // Gaming is used a lot but CFD decided it would be a good brand name to include Gaming...
	"SSD",    // Everything here is an SSD. Only screws up a Gigabyte one since the model is just SSD...
	"PRO",    // Pro in titles causes problems
)

func argMatchesDelimiters(r rune) bool {
	return r == ' ' || r == '(' || r == ')' || r == '/' || r == '_'
}

func prepareProcessedData() map[modelKey][]string {
	file := readCsvFile()
	modelMap := make(map[modelKey][]string)
	for i, row := range file {
		if i == 0 {
			// Skip the headers
			continue
		}
		brand := strings.ToUpper(row[0])
		model := strings.ToUpper(row[1])
		productPages := strings.ToUpper(strings.Join([]string{row[15], row[16]}, " ")) // 15,16 are product pages
		others := strings.ToUpper(strings.Join(row[0:17], " "))

		brandSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(brand, argMatchesDelimiters) {
			brandSet.Add(token)
		}
		modelSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(model, argMatchesDelimiters) {
			modelSet.Add(token)
		}
		productPagesSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(productPages, argMatchesDelimiters) {
			modelSet.Add(token)
		}
		othersSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(others, argMatchesDelimiters) {
			othersSet.Add(token)
		}

		key := modelKey{
			brand:        brandSet.Difference(bannedWords),
			model:        modelSet.Difference(bannedWords),
			productPages: productPagesSet, // We don't ignore them for product pages to help break ties (hopefully)
			others:       othersSet,
		}
		if len(key.brand.ToSlice()) == 0 {
			continue
		}
		if len(key.brand.ToSlice()) == 0 {
			fmt.Println("Found a row with no found model", key.model)
			continue
		}

		modelMap[key] = row
		fmt.Printf("\tBrand: %s : %s : %s\n", key.brand, key.model, row)
	}

	return modelMap
}
