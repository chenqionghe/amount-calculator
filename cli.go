package amountcalculator

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandParams struct {
	Max      float64
	Amounts  []float64
	Overflow float64
}

func parseCliParams() *CommandParams {

	var Max float64
	var Amounts string
	var Overflow float64
	flag.Float64Var(&Max, "max", 0, `target amount, for example:2000`)
	flag.Float64Var(&Overflow, "overflow", 0, `amount allow to overflow, for example: 1`)
	flag.StringVar(&Amounts, "amounts", "", `all your amounts, for example:1,12,123,23,234`)

	flag.Parse()

	if Max == 0 || Amounts == "" {
		flag.Usage()
		os.Exit(1)
	}

	if Max == 0 {
		fmt.Println("max")
		os.Exit(1)
	}
	if Amounts == "" {
		fmt.Println("请输出所有的发票")
		os.Exit(1)
	}

	AmountsFloat := make([]float64, 0)
	data := strings.Split(Amounts, ",")
	for _, v := range data {
		d, _ := strconv.ParseFloat(v, 10)
		AmountsFloat = append(AmountsFloat, d)
	}
	return &CommandParams{
		Overflow: Overflow,
		Max:      Max,
		Amounts:  AmountsFloat,
	}
}

// items:所有发票
// maxValue:目标金额
// overflow:允许误差金额
func CliMode() {
	params := parseCliParams()
	obj := New(params.Amounts, params.Max, params.Overflow)
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
