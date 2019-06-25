import pandas as pd
from datetime import datetime

DataPathMap = {
    "ICD9CM": "../data/icd9cm.csv",
    "ICD9PCS": "../data/icd9pcs.csv",
    "ICD10CM": "../data/icd10cm.csv",
    "ICD10PCS": "../data/icd10pcs.csv"
}

OutputFile = "../convertor/ccs_codes.go"
outputDict = {}

outputTemplate = f'''
package convertor

// ******************************************
// Generated Code : DO NOT MODIFY
// Date : {datetime.now()}
// ******************************************

const CcsCodesJson = `
%s
`
'''


def clean(s):
    s = str.strip(s)
    s = s.replace("'", "")
    s = s.replace("`", "'")
    s = str.strip(s)
    return s


def trimQuote(row):
    return row.apply(clean)


if __name__ == '__main__':
    for codeSetName, filePath in DataPathMap.items():
        print(codeSetName)
        df = pd.read_csv(
            filePath,
            quotechar='"'
        )
        df = df.apply(trimQuote)
        df["icd_code_2"] = df["icd_code"]
        df = df.set_index("icd_code")
        df.rename(columns={'icd_code_2': 'icd_code'}, inplace=True)
        outputDict[codeSetName] = df.to_dict(orient="index")
    import json
    s = json.dumps(outputDict, indent=4)
    with open(OutputFile, 'w') as outfile:
        outfile.write(outputTemplate % s)
