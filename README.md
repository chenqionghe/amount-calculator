# Amount Calculator
High-performance tool for calculating all combinations of matching target amount from the list of given amounts


# Install

```
go get -u github.com/chenqionghe/amount-calculator
```


# Usage
##  Code Example

```go
package main

import (
	"fmt"
	"github.com/chenqionghe/amount-calculator"
)

func main() {
	// all your amounts
	items := []float64{12, 135, 11, 12, 15, 16, 18, 32, 64, 76, 50}
	// target amount
	target := float64(156)
	// allow to overflow
	overflow := float64(1)

	obj := amountcalculator.New(items, target, overflow)
	fmt.Println(obj.GetCombinations())
}

``` 
output
```
[[11 15 16 18 32 64] [16 64 76] [12 18 76 50]]
```





## Commandline Mode Example
create your own go file: main.go
```go
package main

import (
	"github.com/chenqionghe/amount-calculator"
)

func main() {
	amountcalculator.RunCliMode()
}
```
compile
```
go build -o amount-calculator
```
run and pass three parameters
* items: all your amounts
* max: target amount
* overflow: allow to overflow
```shell
 ./amount-calculator -max=156 -overflow=1 -items=12,135,11,12,15,16,18,32,64,76,50
156 [11 15 16 18 32 64]
156 [16 64 76]
156 [12 18 76 50]
157 [12 15 16 18 32 64]
157 [15 16 18 32 76]
157 [15 16 76 50]
156 [11 15 16 18 32 64]
156 [16 64 76]
156 [12 18 76 50]
157 [12 15 16 18 32 64]
157 [15 16 18 32 76]
157 [15 16 76 50]
```