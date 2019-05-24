package main

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type (
	CodeSystem   string
	IcdCssRelMap map[string]IcdCssRel
)

const (
	CodeSystemICD9Proc  CodeSystem = "ICD9PCS"
	CodeSystemICD10Proc CodeSystem = "ICD10PCS"
	CodeSystemICD9Diag  CodeSystem = "ICD9CM"
	CodeSystemICD10Diag CodeSystem = "ICD10CM"
	DATAPATH_icd9cm                = "./data/icd9cm.csv"
	DATAPATH_icd9pcs               = "./data/icd9pcs.csv"
	DATAPATH_icd10cm               = "./data/icd10cm.csv"
	DATAPATH_icd10pcs              = "./data/icd10pcs.csv"
)

var (
	CcsConvertor        map[CodeSystem]IcdCssRelMap
	availableCodeSystem = [4]CodeSystem{
		CodeSystemICD9Proc,
		CodeSystemICD10Proc,
		CodeSystemICD9Diag,
		CodeSystemICD10Diag,
	}
)

type IcdCssRel struct {
	CcsCategory            string
	CcsCategoryDescription string
	IcdCode                string
	IcdDescription         string
}

func init() {
	CcsConvertor = LoadConvertor()
}
func processQuotedString(s string) string {
	return strings.TrimSpace(strings.Trim(s, "'"))
}

func getIcdCssRel(icdCode string, codeSystem CodeSystem) IcdCssRel {
	icdCode = strings.Replace(icdCode, ".", "", 1)
	return CcsConvertor[codeSystem][icdCode]
}

func getIcdCssRelBestEffort(icdCode string) IcdCssRel {
	/*
		Should be avoid for the sake of performance
	*/
	icdCode = strings.Replace(icdCode, ".", "", 1)
	for _, cs := range availableCodeSystem {
		if val, ok := CcsConvertor[cs]; ok {
			if ret, ok := val[icdCode]; ok {
				return ret
			}
		}
	}
	return IcdCssRel{}
}

func loadCsvFileToRecords(filePath string) (*csv.Reader, error) {
	csvfile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(csvfile)
	return r, nil
}

func icdConvertor(filePath string) IcdCssRelMap {
	icdPcs := make(IcdCssRelMap)
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
		icdPcs[processQuotedString(record[0])] = IcdCssRel{
			CcsCategory:            processQuotedString(record[1]),
			CcsCategoryDescription: processQuotedString(record[2]),
			IcdCode:                processQuotedString(record[0]),
			IcdDescription:         processQuotedString(record[3]),
		}
	}
	return icdPcs
}

func LoadConvertor() map[CodeSystem]IcdCssRelMap {
	/*
		NOTE: The oroginal data files have different columns order. I converted them all to become like ICD9 headers.
		* ICD9 headers:
			'ICD-9-CM CODE','CCS CATEGORY','CCS CATEGORY DESCRIPTION','ICD-9-CM CODE DESCRIPTION'

		* ICD10 headers:
			'ICD-10-CM CODE','CCS CATEGORY','ICD-10-CM CODE DESCRIPTION','CCS CATEGORY DESCRIPTION'
	*/

	convertor := make(map[CodeSystem]IcdCssRelMap)
	convertor[CodeSystemICD9Diag] = icdConvertor(DATAPATH_icd9cm)
	convertor[CodeSystemICD9Proc] = icdConvertor(DATAPATH_icd9pcs)
	convertor[CodeSystemICD10Proc] = icdConvertor(DATAPATH_icd10pcs)
	convertor[CodeSystemICD10Diag] = icdConvertor(DATAPATH_icd10cm)

	return convertor
}

func main() {
	fmt.Println(getIcdCssRel("01175", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("618.5", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("518.81", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("V20.2", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("999.4", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("759.0", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("493.00", CodeSystemICD9Diag).CcsCategory)
	fmt.Println(getIcdCssRel("", CodeSystemICD9Diag).CcsCategory)

	fmt.Println(getIcdCssRel("85.42", CodeSystemICD9Proc).CcsCategory)
	fmt.Println(getIcdCssRel("88.98", CodeSystemICD9Proc).CcsCategory)

	fmt.Println(getIcdCssRel("J45.20", CodeSystemICD10Diag).CcsCategory)
	fmt.Println(getIcdCssRel("J45.21", CodeSystemICD10Diag).CcsCategory)

	fmt.Println(getIcdCssRel("009T00Z", CodeSystemICD10Proc).CcsCategory)
	fmt.Println(getIcdCssRel("0DV60CZ", CodeSystemICD10Proc).CcsCategory)

	fmt.Println(getIcdCssRelBestEffort("009T00Z").CcsCategory)
	fmt.Println(getIcdCssRelBestEffort("0DV60CZ").CcsCategory)

	// Not a valid code -> return empty string
	fmt.Println(getIcdCssRelBestEffort("0DV60CZZZZZZZZ").CcsCategory)
}
