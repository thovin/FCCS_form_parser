// package familySheetPopulator
package main

import (
	"encoding/csv"
	"local/commitmentParser"
	"log"
	"os"
	"sort"
)

var INFILE = "C:/Temp/FCCS Re-commitment Form (Responses) - Form Responses 1.csv" //TODO
var OUTPATH = "C:/Temp"                                                           //TODO
var OUTFILENAME = "family_roster.csv"

func main() {
	writeCSV(getFormInput())
}

func writeCSV(input map[string]map[string]string, keys []string) {
	outFile, err := os.Create(OUTPATH + "/" + OUTFILENAME)
	if err != nil {
		log.Println(err)
	}
	w := csv.NewWriter(outFile)
	sort.Strings(keys)
	falsey := ""

	for _, key := range keys {
		data := make([]string, 41)
		row := input[key]

		//TODO I hate that this is brittle, but I don't know a better way
		data[0] = row["applicantLast"]
		data[1] = row["applicantFirst"]
		data[2] = row["spouse"] //TODO did not parse spouse name in input, change?
		data[3] = falsey        //enrollment status
		data[4] = row["volunteerAreas"]
		data[5] = row["preferredPayment"]
		data[6] = row["email"]
		data[7] = row["phoneNumber"]
		data[8] = row["spouseNumber"]
		data[9] = row["s1First"] + " " + row["s1Last"]
		data[10] = row["s1grade"]
		data[11] = row["s1health"]
		data[12] = row["s2First"] + " " + row["s2Last"]
		data[13] = row["s2grade"]
		data[14] = row["s2health"]
		data[15] = row["s3First"] + " " + row["s3Last"]
		data[16] = row["s3grade"]
		data[17] = row["s3health"]
		data[18] = row["s4First"] + " " + row["s4Last"]
		data[19] = row["s4grade"]
		data[20] = row["s4health"]
		data[21] = "" //5th child on families sheet, not on input form
		data[22] = "" //^
		data[23] = "" //^
		data[24] = row["streetAddress"]
		data[25] = row["city"]
		data[26] = row["state"]
		data[27] = row["zip"]
		data[28] = row["church"]
		data[29] = falsey //school district
		data[30] = falsey //records request
		data[31] = falsey //records received
		data[32] = falsey //church school enrollment sent
		data[33] = falsey //withdrawal date
		data[34] = falsey //1st quarter
		data[35] = falsey //2nd quarter
		data[36] = falsey //3rd quarter
		data[37] = falsey //4th quarter
		data[38] = falsey //attendance
		data[39] = row["faithResponse"]
		data[40] = row["policyResponse"]

		if err := w.Write(data); err != nil {
			log.Println(err)
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		log.Println("write error")
	}

	if err := outFile.Close(); err != nil {
		log.Println("error closing output file lock")
	}

}

func getFormInput() (data map[string]map[string]string, keys []string) {
	return commitmentParser.GetParsedCSV(INFILE)
}
