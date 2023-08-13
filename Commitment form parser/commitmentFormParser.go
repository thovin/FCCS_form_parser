package commitmentParser

import {
	"encoding/csv"
	"log"
}

func GetParsedCSV(string filepath) []map[string]string {
	return parseCSV(filepath)
}

func parseCSV (string filepath) []map[string]string{
	output []map[string]string := make([]map[string]string)	//TODO am I doing this slice right?
	inFile, err := os.Open("records.csv")
	if err != nil {
		log.Println(err)
	}
	r := csv.NewReader(inFile)	//TODO add header functionality?

	for {
		in, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
		}

		temp := make(map[string]string)

		temp["timestamp"] == in[0]
		temp["email"] == in[1]
		temp["faithResponse"] == in[2]
		temp["policyResponse"] == in[3]
		temp["volunteerAreas"] == in[4]
		parsedName := in[5]
		temp["applicantFirst"] = parsedName[0]
		temp["applicantMiddle"] = parsedName[1]
		temp["applicantLast"] = parsedName[2]
		temp["streetAddress"] == in[6]
		temp["city"] == in[7]
		temp["state"] == in[8]
		temp["zip"] == in[9]
		temp["phoneNumber"] == in[10]
		temp["Employer"] == in[11]
		temp["workPhone"] == in[12]
		temp["workSchedule"] == in[13]
		temp["church"] == in[14]
		temp["spouse"] == in[15]
		temp["spouseEmployer"] == in[16]
		temp["spouseNumber"] == in[17]
		temp["spouseEmail"] == in[18]
		temp["preferredPayment"] == in[19]
		parsedName = in[20]
		temp["s1First"] = parsedName[0]
		temp["s1Middle"] = parsedName[1]
		temp["s1Last"] = parsedName[2]
		temp["s1grade"] == in[21]
		temp["s1health"] == in[22]
		temp["addon1"] == in[23]
		parsedName = in[24]
		temp["s2First"] = parsedName[0]
		temp["s2Middle"] = parsedName[1]
		temp["s2Last"] = parsedName[2]
		temp["s2grade"] == in[25]
		temp["s2health"] == in[26]
		temp["addon2"] == in[27]
		parsedName = in[28]
		temp["s3First"] = parsedName[0]
		temp["s3Middle"] = parsedName[1]
		temp["s3Last"] = parsedName[2]
		temp["s3grade"] == in[29]
		temp["s3health"] == in[30]
		temp["addon3"] == in[31]
		parsedName = in[32]
		temp["s4First"] = parsedName[0]
		temp["s4Middle"] = parsedName[1]
		temp["s4Last"] = parsedName[2]
		temp["s4grade"] == in[33]
		temp["s4health"] == in[34]
		temp["vaccinationStatus"] == in[35]
		//column AK blank
		temp["moneyInSpreadsheet"] == in[37]

		output = append(output, temp)

	}
	
	return output
}

func splitName(string in) []string {
	output = make([]string, 3)	//TODO syntax?

	temp := strings.Split(in)
	output[0] = temp[0]

	if len(temp) == 3 {
		output[1] == temp[1]
		output[2] == temp[2]
	} else if len(temp) == 2 {
		output[2] == temp[1]
	} else if len(temp) == 0 || len(temp) == 1{
		log.Println("Name empty or singular")
	} else if len(temp) > 3 {
		log.Println("Name has 4 or more fields")
	}

	return output	//TODO is default middle name correctly empty string?
}