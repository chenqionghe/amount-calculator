# Amount Calculator
High performance tools for calculating the amounts from the list of given amounts


# Install

```
go get -u github.com/chenqionghe/amount-calculator
```


# Usage

```
package main

import (
	"fmt"
	"github.com/chenqionghe/amount-calculator"
)

func main() {
	items := []float64{12, 135, 11, 100, 12, 15, 16, 18, 32, 64, 76, 50}
	target := float64(156)
	overflow := float64(1)
	obj := amountcalculator.New(items, target, overflow)
	fmt.Println(obj.Run())
}

``` 






* Run Commandline

```

create your own go file: main.go
```
func main() {
	RunCliMode()
}
```

run this file, for example
```
go run main -max=156 -overflow=1 -amounts=12,135,11,100,12,15,16,18,32,64,76,50

```