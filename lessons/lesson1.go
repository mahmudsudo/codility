package lessons

import (
	"strconv"
	"strings"
)

func GetBinaryGap(N int) int{

	bin := strconv.FormatInt(int64(N),2)
	binArr:= strings.Split(bin,"")
	if len(binArr) == 0{
		return 0
	}
	// var gapCount int =0
	var gapLen int = 0
	if binArr[0] =="1" && binArr[1] == "0"{
		gapLen++
		for i,r := range binArr[1:]{
		// if gapCount > 1{

		// }
		if r == "1"{
			if binArr[i+1] == "0"{
				gapLen++
			}
			// gapCount++
			

		}else if r == "0" && binArr[i+1] =="0"{

			gapLen++
			if binArr[len(binArr) -1] != "1"{
				gapLen = 0
			}
		}else{
			gapLen = 0
		}
	}
	}
	
	return gapLen
}

var Me string =" mahmud"