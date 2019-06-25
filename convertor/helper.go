package convertor

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

const (
	DATAPATH_ICD9CM   = "data/icd9cm.csv"
	DATAPATH_ICD9PCS  = "data/icd9pcs.csv"
	DATAPATH_ICD10CM  = "data/icd10cm.csv"
	DATAPATH_ICD10PCS = "data/icd10pcs.csv"
)

func loadCodeSetFromFile() map[CodeSystem]IcdInfoMap {
	convertor := make(map[CodeSystem]IcdInfoMap)
	convertor[CodeSystemICD9Diag] = icdConvertor(DATAPATH_ICD9CM)
	convertor[CodeSystemICD9Proc] = icdConvertor(DATAPATH_ICD9PCS)
	convertor[CodeSystemICD10Proc] = icdConvertor(DATAPATH_ICD10PCS)
	convertor[CodeSystemICD10Diag] = icdConvertor(DATAPATH_ICD10CM)

	return convertor
}

func icdConvertor(filePath string) IcdInfoMap {
	icdPcs := make(IcdInfoMap)
	r, err := loadCsvFileToRecords(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
		return nil
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		icdPcs[processQuotedString(record[0])] = IcdInfo{
			CcsCategory:            processQuotedString(record[1]),
			CcsCategoryDescription: processQuotedString(record[2]),
			IcdCode:                processQuotedString(record[0]),
			IcdDescription:         processQuotedString(record[3]),
		}
	}
	return icdPcs
}

func loadCsvFileToRecords(filePath string) (*csv.Reader, error) {
	csvfile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(csvfile)
	return r, nil
}

func processQuotedString(s string) string {
	return strings.TrimSpace(strings.Trim(s, "'"))
}
