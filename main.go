package main

import (
	"fmt"
	"strconv"
)

func main() {

	// read Portfolio Performance export
	obj, err := loadJSON("Regionen_(MSCI).json")
	if err != nil {
		panic(err)
	}

	//-----------------------------------------------------------

	for ii, instrument := range obj.Instruments {
		if len(instrument.Categories) <= 1 {
			continue // skip single country
		}

		name := instrument.Identifiers.Name
		isin := instrument.Identifiers.ISIN
		fmt.Println(isin, name)

		text := finanzfluss(isin)
		cw := countryWeighting(text)

		for ci, category := range instrument.Categories {
			// portfolio performance data
			country := category.Path[len(category.Path)-1]
			weight := category.Weight

			// new data
			newWeight, ok := cw[country]
			cw[country] = "OK" // mark as used
			newWeightF, _ := strconv.ParseFloat(newWeight, 64)
			obj.Instruments[ii].Categories[ci].Weight = newWeightF // set data

			// print
			if !ok {
				fmt.Printf("err: Country from Portfolio Performance not included in the Finanzfluss dataset: %s = %.2f\n", country, weight)
			} else {
				weightS := fmt.Sprintf("%.2f", weight)
				if weightS == newWeight {
					fmt.Printf("= %s: %s\n", country, weightS)
				} else {
					fmt.Printf("+ %s: %s -> %s  !!!\n", country, weightS, newWeight)
				}
			}
		}

		// find unused countries
		for k, v := range cw {
			if v != "OK" {
				fmt.Printf("unused country: %s\n", k)
			}
		}

		fmt.Println("------------------")
	}

	//-----------------------------------------------------------

	// write new import file
	err = writeJSON("new.json", obj)
	if err != nil {
		panic(err)
	}

}
