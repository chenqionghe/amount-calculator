package amountcalculator

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
)

type AmountCalculator struct {
	maxValue int   //target amount
	items    []int //all the amounts
	overflow int   //amount allow to overflow
}

// items: all your amounts
// maxValue: target amount
// overflow: allow to overflow
func New(items []float64, maxValue float64, overflow float64) *AmountCalculator {
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

func (this *AmountCalculator) dollarToCent(value float64) int {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)

	decimalValue := decimal.NewFromFloat(value)
	decimalValue = decimalValue.Mul(decimal.NewFromInt(100))

	res, _ := decimalValue.Float64()
	return int(res)

}

func (this *AmountCalculator) centToDollar(v int) float64 {
	value := float64(v)
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value/100), 64)
	return value

}

//run and return all combinations
func (this *AmountCalculator) GetCombinations() [][]float64 {
	items := this.items
	n := len(this.items)
	max := this.maxValue + this.overflow
	states := this.createStates(len(this.items), max+1)

	states[0][0] = true
	if items[0] <= max {
		states[0][items[0]] = true
	}

	for i := 1; i < n; i++ {
		//do not select
		for j := 0; j <= max; j++ {
			if states[i-1][j] {
				states[i][j] = states[i-1][j]
			}
		}
		//select
		for j := 0; j <= max-items[i]; j++ {
			if states[i-1][j] {
				states[i][j+items[i]] = true
			}
		}
	}
	//get all combination matching the target
	res := make([][]float64, 0)
	for j := this.maxValue; j <= max; j++ {
		for i := 0; i < n; i++ {
			if states[i][j] {
				//the judgment must be selected last, otherwise there will be a repeated combination
				if i == 0 {
					//the first element has been matched
					res = append(res, this.getSelected(states, items, i, j))
				} else if j-items[i] >= 0 && states[i-1][j-items[i]] == true {
					res = append(res, this.getSelected(states, items, i, j))
				}
			}
		}
	}
	return res

}

//find all selected items
func (this *AmountCalculator) getSelected(states [][]bool, items []int, n, max int) []float64 {
	var selected = make([]int, 0)
	for i := n; i >= 1; i-- {
		//selected
		if max-items[i] >= 0 && states[i-1][max-items[i]] == true {
			selected = append([]int{items[i]}, selected...)
			max = max - items[i]
		} else {
			//do not selected，max unchanged, continue
		}
	}

	//if max is equal to 0, add the first element
	if max != 0 {
		selected = append([]int{items[0]}, selected...)
	}

	dollarItems := make([]float64, len(selected))
	for i, v := range selected {
		dollarItems[i] = this.centToDollar(v)
	}
	return dollarItems
}

//init all states
func (this *AmountCalculator) createStates(n, max int) [][]bool {
	states := make([][]bool, n)
	for i, _ := range states {
		states[i] = make([]bool, max)
	}
	return states
}
