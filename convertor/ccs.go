package convertor

import (
	"encoding/json"
	"strings"
)

type (
	CodeSystem string
	// IcdInfo is extracted from a row in datasource file
	IcdInfo struct {
		CcsCategory            string `json:"ccs_cat"`
		CcsCategoryDescription string `json:"ccs_cat_desc"`
		IcdCode                string `json:"icd_code"`
		IcdDescription         string `json:"icd_desc"`
	}
	// IcdInfoMap is a hash table. The module is used to lookup the CCS of an ICD code, therefore, the IcdCode field is selected as a key in order to fasten lookup.
	/*
		e.g: icdInfoMap = {
			"ICD9CM": "....",
			"ICD9PCS": "....",
			"ICD10PCS": "....",
		}
	*/
	IcdInfoMap map[string]IcdInfo
	// CcsConvertorType is used to categorize icd codes in to categories for faster lookup
	CcsConvertorType map[CodeSystem]IcdInfoMap
)

const (
	CodeSystemICD9Proc  CodeSystem = "ICD9PCS"
	CodeSystemICD10Proc CodeSystem = "ICD10PCS"
	CodeSystemICD9Diag  CodeSystem = "ICD9CM"
	CodeSystemICD10Diag CodeSystem = "ICD10CM"
)

var (
	// CcsConvertor is going to be initialized when the mofule is loaded. All the lookup will be performed via this variable.
	CcsConvertor CcsConvertorType

	availableCodeSystem = [4]CodeSystem{
		CodeSystemICD9Proc,
		CodeSystemICD10Proc,
		CodeSystemICD9Diag,
		CodeSystemICD10Diag,
	}
)

func init() {
	var err error
	CcsConvertor, err = loadCodeSetFromString(CcsCodesJson)
	if err != nil {
		panic("[E0001] Cannot load CCS convertor. Please review the CcsCodesJson string.")
	}
}

func loadCodeSetFromString(s string) (CcsConvertorType, error) {
	r := strings.NewReader(s)
	var cb CcsConvertorType
	dec := json.NewDecoder(r)
	err := dec.Decode(&cb)
	if err != nil {
		return CcsConvertorType{}, err
	}

	return cb, nil
}

func GetIcdInfoBestEffort(icdCode string) IcdInfo {
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
	return IcdInfo{}
}

func GetIcdCssRel(icdCode string, codeSystem CodeSystem) IcdInfo {
	icdCode = strings.Replace(icdCode, ".", "", 1)
	return CcsConvertor[codeSystem][icdCode]
}
