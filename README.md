# ICD to CCS convertor 

## About the mapping
ICD have 2 versions ICD9 and ICD10. Each version have 2 kinds of code CM (diagnosis) and PCS (procedure). 
Thus, we have 4 combinations.

The mapping can be found at [ICD9 to CCS](https://www.hcup-us.ahrq.gov/toolssoftware/ccs/ccs.jsp) and [ICD10 to CCS](https://www.hcup-us.ahrq.gov/toolssoftware/ccs10/ccs10.jsp)

Note: the mapping for ICD10 is still in beta mode. Thus, we need to update it arcordingly.

## Usage
* Access via the variable CcsConvertor

## Development
The original data files have different columns order. I converted them all to become like ICD9 headers	

| Version | Original header                                                                         |   |
|---------|-----------------------------------------------------------------------------------------|---|
| ICD9    | 'ICD-9-CM CODE','CCS CATEGORY','CCS CATEGORY DESCRIPTION','ICD-9-CM CODE DESCRIPTION'   |   |
| ICD10   | 'ICD-10-CM CODE','CCS CATEGORY','ICD-10-CM CODE DESCRIPTION','CCS CATEGORY DESCRIPTION' |   |