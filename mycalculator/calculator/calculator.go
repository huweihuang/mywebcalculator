package calculator

import (
	"fmt"
	"math"
	"strconv"

	"mywebcalculator/mycalculator/data"
	"mywebcalculator/mycalculator/functions"
)

type Calculator interface {
	//运行计算器
	Run() float64
	//基本运算规则
	BasicCalculation(a, b float64, operation string) float64
	//计算运算结果
	Calculate(datas []string) float64
}

type calculator struct {
	//运算表达式
	operationExpression string
	//运算结果
	result float64
}

func NewCalculator(operationExpression string) *calculator {

	ca := calculator{
		operationExpression: operationExpression,
	}

	return &ca
}

func (ca *calculator) Run() float64 {

	var arr []string = functions.GetPostfixExpression(ca.operationExpression)

	ca.result = ca.Calculate(arr)

	return ca.result

}

//定义基本运算法则
func (ca *calculator) BasicCalculation(a, b float64, operation string) (result float64,err error){
	switch operation {
	case "+":
		result=a + b
	case "-":
		result=a - b
	case "*":
		result=a * b
	case "/":
		if b == 0 {
			return a,fmt.Errorf("The dividend can not be zero: %v", b)
		} else {
			result=a / b
		}
	case "^":
		result=math.Pow(a, b)
	default:
		return -1,fmt.Errorf("Illegal input！ %v,%v", a,b)
	}
	return result,nil
}

//由后缀表达式计算结果
func (ca *calculator) Calculate(datas []string) float64 {
	var stack data.LinkStack
	stack.Init()
	for i := 0; i < len(datas); i++ {
		if functions.IsNumberString(datas[i]) {
			if f, err := strconv.ParseFloat(datas[i], 64); err != nil {
				panic("operatin process go wrong.")
			} else {
				stack.Push(f)
			}
		} else {
			p1 := stack.Pop().(float64)
			p2 := stack.Pop().(float64)
			p3,err:= ca.BasicCalculation(p2, p1, datas[i])
			if err!=nil{
				fmt.Println(err)
			}
			stack.Push(p3)
		}
	}
	res := stack.Pop().(float64)
	return res
}
