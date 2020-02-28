package main

import "fmt"

//items:所有发票 maxValue:目标金额 overflow:允许误差金额
func RunByComandParams() {
	//parse command params
	params := ParseParams()
	obj := NewAmountCalculator(params.Amounts, params.Max, params.Overflow)
	//执行计算，返回所有结果方案
	res := obj.Run()
	//打印所有方案
	for _, v := range res {
		fmt.Print(sum(v), " ")
		fmt.Println(v)
	}
}

func sum(a []float64) float64 {
	var s float64 = 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
