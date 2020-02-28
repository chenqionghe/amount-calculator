package amountcalculator

import (
	"flag"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
)

type AmountCalculator struct {
	maxValue int   //期望值（单元为分）
	items    []int //发票金额（单元为分）
	overflow int   //允许的误差值（单元为分）
}

//items:所有发票 maxValue:目标金额 overflow:允许误差金额
func NewAmountCalculator(items []float64, maxValue float64, overflow float64) *AmountCalculator {
	obj := &AmountCalculator{}
	obj.maxValue = obj.dollarToCent(maxValue)
	obj.overflow = obj.dollarToCent(overflow)
	centItems := make([]int, len(items))
	for i, v := range items {
		centItems[i] = obj.dollarToCent(v)
	}
	obj.items = centItems
	return obj
}

//元转分
func (this *AmountCalculator) dollarToCent(value float64) int {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)

	decimalValue := decimal.NewFromFloat(value)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

	res, _ := decimalValue.Float64()
	return int(res)

}

//分转元
func (this *AmountCalculator) centToDollar(v int) float64 {
	value := float64(v)
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value/100), 64)
	return value

}

//执行计算，返回所有方案
func (this *AmountCalculator) Run() [][]float64 {
	items := this.items
	n := len(this.items)
	max := this.maxValue + this.overflow
	states := this.createStates(len(this.items), max+1)

	states[0][0] = true
	if items[0] <= max {
		states[0][items[0]] = true
	}

	for i := 1; i < n; i++ {
		//不选
		for j := 0; j <= max; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}
		//选中
		for j := 0; j <= max-items[i]; j++ {
			if states[i-1][j] {
				states[i][j+items[i]] = true
			}
		}
	}
	//获取最终所有满足的方案
	res := make([][]float64, 0)
	for j := this.maxValue; j <= max; j++ {
		for i := 0; i < n; i++ {
			if states[i][j] {
				//判断必须最后一个选中才算，要不区间有重合 比如前5个元素已经满足目标金额了，state[5][w]=true，然后state[6][w]也是true，存在重复的方案
				if i == 0 {
					//第一个元素已经满足
					res = append(res, this.getSelected(states, items, i, j))
				} else if j-items[i] >= 0 && states[i-1][j-items[i]] == true {
					res = append(res, this.getSelected(states, items, i, j))
				}
			}
		}
	}
	return res

}

//获取所有选中的元素（倒推）
func (this *AmountCalculator) getSelected(states [][]bool, items []int, n, max int) []float64 {
	var selected = make([]int, 0)
	for i := n; i >= 1; i-- {
		//元素被选中
		if max-items[i] >= 0 && states[i-1][max-items[i]] == true {
			selected = append([]int{items[i]}, selected...)
			max = max - items[i]
		} else {
			//没选，max重量不变，直接进入下一次
		}
	}

	//如果max不为0，说明还需要追加第一个元素
	if max != 0 {
		selected = append([]int{items[0]}, selected...)
	}

	dollarItems := make([]float64, len(selected))
	for i, v := range selected {
		dollarItems[i] = this.centToDollar(v)
	}
	return dollarItems
}

//初始化所有状态
func (this *AmountCalculator) createStates(n, max int) [][]bool {
	states := make([][]bool, n)
	for i, _ := range states {
		states[i] = make([]bool, max)
	}
	return states
}

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
