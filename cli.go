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
	items    []float64
	Overflow float64
}

func parseCliParams() *CommandParams {

	var Max float64
	var items string
	var Overflow float64
	flag.Float64Var(&Max, "max", 0, `target amount, for example:2000`)
	flag.Float64Var(&Overflow, "overflow", 0, `amount allow to overflow, for example: 1`)
	flag.StringVar(&items, "items", "", `all your amounts, for example:1,12,123,23,234`)

	flag.Parse()

	if Max == 0 || items == "" {
		flag.Usage()
		os.Exit(1)
	}

	if Max == 0 {
		fmt.Println("max")
		os.Exit(1)
	}
	if items == "" {
		fmt.Println("请输出所有的发票")
		os.Exit(1)
	}

	itemsFloat := make([]float64, 0)
	data := strings.Split(items, ",")
	for _, v := range data {
		d, _ := strconv.ParseFloat(v, 10)
		itemsFloat = append(itemsFloat, d)
	}
	return &CommandParams{
		Overflow: Overflow,
		Max:      Max,
		items:    itemsFloat,
	}
}

func RunCliMode() {
	params := parseCliParams()
	obj := New(params.items, params.Max, params.Overflow)
	res := obj.GetCombinations()
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
