package main

import (
	"fmt"

	"github.com/zquangu112z/IcdCcs/convertor"
)

func main() {
	fmt.Println(convertor.GetIcdInfoBestEffort("C028").CcsCategory)
	fmt.Println(convertor.GetIcdInfoBestEffort("C028").CcsCategoryDescription)

	fmt.Println(convertor.GetIcdInfoBestEffort("Z44001").CcsCategory)
	fmt.Println(convertor.GetIcdInfoBestEffort("Z44001").CcsCategoryDescription)
}
