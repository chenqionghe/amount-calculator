package lib

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

func ParseParams() *CommandParams {

	var Max float64
	var Amounts string
	var Overflow float64
	flag.Float64Var(&Max, "max", 0, `目标最大金额，例如：2000`)
	flag.Float64Var(&Overflow, "overflow", 0, `允许超出的金额，例如：1`)

	flag.StringVar(&Amounts, "amounts", "", `所有的碎票，多个用,隔开,例如：280,280,280,280,280,280,250,250,250,230,220,215`)

	flag.Parse()

	if Max == 0 || Amounts == "" {
		flag.Usage()
		os.Exit(1)
	}

	if Max == 0 {
		fmt.Println("请输出最大金额")
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
