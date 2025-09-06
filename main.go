package main

import (
	"fmt"
	"strconv"
)

func main() {

	// read Portfolio Performance export
	obj, err := loadJSON("Regionen_(MSCI)_(ex.A).json")
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
			newWeightF, _ := strconv.ParseFloat(newWeight, 64)
			obj.Instruments[ii].Categories[ci].Weight = newWeightF // set data

			// print
			if !ok {
				fmt.Printf("err: Country from Portfolio Performance not included in the Finanzfluss dataset: %s = %.2f\n", country, weight)
			} else {
				fmt.Printf("%s: %.2f -> %s\n", country, weight, newWeight)
			}
		}
		fmt.Println("------------------")
	}

	//-----------------------------------------------------------

	// write new import file
	err = writeJSON("Regionen_(MSCI)_(ex.A)-2.json", obj)
	if err != nil {
		panic(err)
	}
	
}
