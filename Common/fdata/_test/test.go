package main

import (
	//"io"
	//"bufio"
	"fmt"
	//"os"
	"rcdd/project/infoBase/Common/fdata"
	//"strings"
)

var (
	pp = fmt.Println
)

func testRead(sdfile string) {
	mols, err := fdata.ReadByFile(sdfile, 0)
	if err != nil {
		panic(err)
	}

	pp(mols.ToString())
}

func main() {
	testRead("E:/iprexmed/product/infoBase/testData/debug/Specs_NP_1mg_201703-for-infobase-2.sdf")
}
