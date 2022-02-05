package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("110010", "100"))
}

func addBinary(a string, b string) string {
	res, sum, aDx, bDx := "", 0, len(a)-1, len(b)-1

	for aDx >= 0 || bDx >= 0 || sum/2 > 0 {
		sum = (sum / 2) + getBint(a, aDx) + getBint(b, bDx)
		res = fmt.Sprintf("%d%s", sum%2, res)
		aDx--
		bDx--
	}
	return res
}

func getBint(s string, i int) int {
	if i < 0 {
		return 0
	}
	return map[string]int{"0": 0, "1": 1}[string(s[i])]
}

//func addBinary(a string, b string) string {
//	rem := 0
//	res := ""
//	aIdx := len(a) - 1
//	bIdx := len(b) - 1
//	var lStr string
//	if aIdx > bIdx {
//		lStr = a
//	} else {
//		lStr = b
//	}
//
//	for aIdx >= 0 && bIdx >= 0 {
//		currA := string(a[aIdx])
//		currB := string(b[bIdx])
//		if currA == "1" && currB == "1" {
//			if rem == 1 {
//				res = "1" + res
//			} else {
//				res = "0" + res
//				rem = 1
//			}
//		} else if currA == "0" && currB == "0" {
//			if rem == 1 {
//				res = "1" + res
//				rem = 0
//			} else {
//				res = "0" + res
//			}
//		} else {
//			if rem == 1 {
//				res = "0" + res
//			} else {
//				res = "1" + res
//			}
//		}
//		aIdx--
//		bIdx--
//	}
//
//	itr := int(math.Max(float64(aIdx), float64(bIdx)))
//
//	if rem == 1 || itr >= 0 {
//		for rem == 1 || itr >= 0 {
//			if itr < 0 || string(lStr[itr]) == "0" {
//				if rem == 1 {
//					res = "1" + res
//					rem = 0
//				} else {
//					res = "0" + res
//				}
//				itr--
//			} else {
//				if rem == 1 {
//					if string(lStr[itr]) == "1" {
//						res = "0" + res
//					} else {
//						res = "1" + res
//						rem = 0
//					}
//				} else {
//					res = string(lStr[itr]) + res
//				}
//
//				itr--
//			}
//		}
//	}
//
//	return res
//}
