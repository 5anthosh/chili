# 🌶️ chili

> Currently in development, Unstable (API may change in future)

Simple expression evaluation engine.

Expression is one liner that evalutes into single value

### Features

- Number accuracy (using github.com/shopspring/decimal pkg)
- Extensible
- Simple grammer

### Installation

```shell
$ go get github.com/5anthosh/chili
```

### Examples

```go
package main

import (
    "fmt"
    "github.com/5anthosh/chili"
)
func main() {
    expression := "val > 50 ? 'Greater than or 50' : 'Smaller than 50'"
    values := map[string]interface{}{
        "val": 60,
    }
    result, err := chili.Eval(expression, values)
    if err != nil {
        panic(err)
    }
    println(fmt.Sprintf("%v result", result))
}
```

```go
package main

import (
    "fmt"
    "github.com/5anthosh/chili/environment"
    "github.com/5anthosh/chili/evaluator"
    "github.com/5anthosh/chili/parser"
    "github.com/shopspring/decimal"
)

func main() {
    source := "PI*R^2 + abs(45.345)"

    env := environment.New()
    env.SetDefaultFunctions()
    env.SetDefaultVariables()
    env.SetIntVariable("R", 2)

    chiliParser := parser.New(source)
    expression, err := chiliParser.Parse()
    if err != nil {
        panic(err)
    }

    chiliEvaluator := evaluator.New(env)
    value, err := chiliEvaluator.Run(expression)
    if err != nil {
        panic(err)
    }

    println(fmt.Sprintf("%v result", value))
}
```

## License

[MIT](https://github.com/5anthosh/chili/blob/main/LICENSE)
