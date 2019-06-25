# Generate ICD to CCS mapping
* The module is used to generate codeset mapping for the convertor from ICD to CCS.
* The module is written in `python`.
## Usage
### Preparation
1. Put datasource file (which is download from CMS) in to `../../data/icd-ccs-mapping/`
2. Make sure the header of these `.csv` files is in order as follow: icd_code, ccs_cat, ccs_cat_desc, icd_desc
### Generate
```Makefile
make run
```

The output file is `../../qm/mssp2019/ccs_codes.go`
