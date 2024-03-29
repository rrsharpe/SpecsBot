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
	"",        // Some product page cells are empty
	"BLACK",   // WD Black models don't always have Black listed in the table
	"RED",     // WD Black models don't always have Red listed in the table
	"NVME",    // Causes false positives across many brands
	"PCIE",    // Causes false positives across many brands
	"4.0",     // Causes false positives across many brands
	"GAMING",  // Gaming is used a lot but CFD decided it would be a good brand name to include Gaming...
	"SSD",     // Everything here is an SSD. Only screws up a Gigabyte one since the model is just SSD...
	"PRO",     // Pro in titles causes problems
	"3D",      // Many things have 3D in the model name
	"SPATIUM", // This is inconsistently included in the spreadsheet for MSI
)

// Common brand aliases
var brandAliases = map[string][]string{
	"TEAM":  {"TEAMGROUP"},
	"ADATA": {"XPG"},
}

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
		validOthers := append([]string{}, row[0:4]...)                                 // 4 is capacity
		validOthers = append(validOthers, row[5])                                      // 6:configuration, 7:dram, 8:hmb
		validOthers = append(validOthers, row[9:13]...)                                // 13: categories
		validOthers = append(validOthers, row[14:17]...)
		others := strings.ToUpper(strings.Join(validOthers, " "))

		brandSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(brand, argMatchesDelimiters) {
			if aliases, isPresent := brandAliases[token]; isPresent {
				for _, alias := range aliases {
					brandSet.Add(alias)
				}
			}
			brandSet.Add(token)
		}
		modelSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(model, argMatchesDelimiters) {
			modelSet.Add(token)
		}
		productPagesSet := mapset.NewSet()
		for _, token := range strings.FieldsFunc(productPages, argMatchesDelimiters) {
			productPagesSet.Add(token)
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
