package familySheetPopulator

import (
	"encoding/csv"
	"local/commitmentParser"
	"log"
	"os"
	"sort"
)

var INFILE = "filepath"  //TODO
var OUTPATH = "filepath" //TODO
var OUTFILENAME = "family_roster.csv"

func Main() {

}

func writeCSV(in []map[string]string) { //TODO merge with old info?
	outFile, err := os.Create(OUTPATH + "/" + OUTFILENAME)
	if err != nil {
		log.Println(err)
	}
	w = csv.NewWriter(outFile)
	input, keys := getFormInput()
	sort.Strings(keys)
	falsey = "FILL ME"

	for _, key := range keys {
		data := make([]string, len(input[row])) //TODO right length since different columns?
		data := make([]string, 34)              //TODO temp
		row := input[key]

		//TODO I hate that this is brittle, but I don't know a better way
		data[0] = row["applicantLast"]
		data[1] = row["applicantFirst"]
		data[2] = row["spouse"]  //TODO did not parse spouse name in input, change?
		data[3] = falsey         //enrollment status
		data[4] = row["s1First"] //first or full name for students?
		data[5] = row["s1grade"]
		data[6] = row["s2First"]
		data[7] = row["s2grade"]
		data[8] = row["s3First"]
		data[9] = row["s3grade"]
		data[10] = row["s4First"]
		data[11] = row["s4grade"]
		data[12] = falsey //5th child on families sheet, not on input form
		data[13] = falsey //^
		data[14] = row["streetAddress"]
		data[15] = row["city"]
		data[16] = row["state"]
		data[17] = row["zip"]
		data[18] = row["phoneNumber"]
		data[19] = row["spouseNumber"]
		data[20] = row["email"]
		data[21] = row["church"]
		data[22] = falsey //school district
		data[23] = falsey //checked for new
		data[24] = falsey //records request
		data[25] = falsey //records recieved
		data[26] = falsey //church school
		data[27] = falsey //notes
		data[28] = falsey //withdrawal date
		data[29] = falsey //1st quarter
		data[30] = falsey //2nd quarter
		data[31] = falsey //3rd quarter
		data[32] = falsey //4th quarter
		data[33] = falsey //attendance
		data[34] = falsey //end of quarter contact w/ first
		// data[35] = row[""]
		// data[36] = row[""]
		// data[37] = row[""]
		// data[38] = row[""]
		// data[39] = row[""]
		// data[40] = row[""]
		// data[41] = row[""]
		// data[42] = row[""]
		// data[43] = row[""]
		// data[44] = row[""]
		// data[45] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]
		// data[] = row[""]

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
	return commitmentParser.GetParsedCSV(INFILE) //TODO am I calling the right thing? Package name or file?
}
