# ICD to CCS code convertor 

## About the mapping
ICD have 2 versions ICD9 and ICD10. Each version have 2 kinds of code CM (diagnosis) and PCS (procedure). 
Thus, we have 4 combinations.

The mapping can be found at [ICD9 to CCS](https://www.hcup-us.ahrq.gov/toolssoftware/ccs/ccs.jsp) and [ICD10 to CCS](https://www.hcup-us.ahrq.gov/toolssoftware/ccs10/ccs10.jsp)

Note: the mapping for ICD10 is still in `beta` mode. Thus, we need to update it arccordingly.

## Usage
Sample is in main.go
### GetIcdInfo()
If we already know the code system of the code, we can use this function.
### GetIcdInfoBestEffort()
Just provide the ICD code. The script will try to retrieve the info.

## Development
The original data mapping files have different columns order. I converted them as follow order: icd_code, ccs_cat, ccs_cat_desc, icd_desc

| Version | Original header                                                                         |   |
|---------|-----------------------------------------------------------------------------------------|---|
| ICD9    | 'ICD-9-CM CODE','CCS CATEGORY','CCS CATEGORY DESCRIPTION','ICD-9-CM CODE DESCRIPTION'   |   |
| ICD10   | 'ICD-10-CM CODE','CCS CATEGORY','ICD-10-CM CODE DESCRIPTION','CCS CATEGORY DESCRIPTION' |   |

### Two way initial the Ccs Convertor
The CcsConvertor variable is loaded at `convertor/ccs.go/init()`:

* Load directly from csv file
```go
CcsConvertor = loadCodeSetFromFile()
```
* Load from json string
```
CcsConvertor, err = loadCodeSetFromString(CcsCodesJson)
```

This way, we need to generate the `CcsCodesJson`, please read `./scripts/README.md`
