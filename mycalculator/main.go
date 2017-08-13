package main

import (
	"bufio"
	"fmt"
	"os"

	"mycalculator/calculator"
)

func main() {

	//打印说明
	fmt.Println(`
        This is a simple calculator!
        Please enter an operation expression.
    `)

	//读取输入信息
	r := bufio.NewReader(os.Stdin)
	//解析输入参数
	for {
		fmt.Print("Enter an operation expression-> ")

		rawLine, _, err := r.ReadLine()

		errorHandler(err)

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		ca := calculator.NewCalculator(line)

		result := ca.Run()

		fmt.Println("The result is：", result)
		fmt.Println()
	}
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
