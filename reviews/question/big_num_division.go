/*
给定一个字符串，内容是数字字符，求该字符串表达的数值除以数值n(n<10)的结果，
除不尽的保留2位小数（给出解题思路，并代码实现）
例如：给定字符串”4562”除以5的结果是”912.4”
函数原型: char* DivString(char *dividend, int divisor);
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	result := DivString("1000000000000000000000000000000000000000000", 11)
	fmt.Println(result)
}

func DivString(char string, divisor int) string {
	result := []float64{}
	resultStr := ""

	if divisor == 0 {
		return ""
	}

	index := 0
	var presentNum float64

	for k, v := range char {
		iNum, _ := strconv.ParseFloat(string(v), 10) // float 转换
		iNum = presentNum + iNum + float64(index)*10
		R := int(iNum) % divisor
		N := math.Floor(iNum / float64(divisor))

		if k == len(char)-1 {
			N = iNum / float64(divisor)
			N = math.Floor(N * 100)
			N = N / 100
		}

		result = append(result, N)
		index = R
	}
	breakStatus := true
	for _, v := range result {
		if v == 0 && breakStatus {
			breakStatus = true
			continue
		} else {
			breakStatus = false
		}
		resultStr = resultStr + strconv.FormatFloat(v, 'g', 12, 64)
	}

	return resultStr
}
