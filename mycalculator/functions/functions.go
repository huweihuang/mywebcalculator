package functions

import (
	"mywebcalculator/mycalculator/data"
)

//由中缀表达式获取后缀表达式
func GetPostfixExpression(exp string) []string {

	var stack data.LinkStack
	stack.Init()

	var spiltedStr []string = convertToStrings(exp)
	var datas []string

	for i := 0; i < len(spiltedStr); i++ { // 遍历每一个字符
		tmp := spiltedStr[i] //当前字符

		if !IsNumberString(tmp) { //是否是数字
			// 四种情况入栈
			// 1 左括号直接入栈
			// 2 栈内为空直接入栈
			// 3 栈顶为左括号，直接入栈
			// 4 当前元素不为右括号时，在比较栈顶元素与当前元素，如果当前元素大，直接入栈。
			if tmp == "(" ||
				stack.LookTop() == nil || stack.LookTop().(string) == "(" ||
				(compareOperator(tmp, stack.LookTop().(string)) == 1 && tmp != ")") {
				stack.Push(tmp)
			} else { // ) priority
				if tmp == ")" { //当前元素为右括号时，提取操作符，直到碰见左括号
					for {
						if pop := stack.Pop().(string); pop == "(" {
							break
						} else {
							datas = append(datas, pop)
						}
					}
				} else { //当前元素为操作符时，不断地与栈顶元素比较直到遇到比自己小的（或者栈空了），然后入栈。
					for {
						pop := stack.LookTop()
						if pop != nil && compareOperator(tmp, pop.(string)) != 1 {
							datas = append(datas, stack.Pop().(string))
						} else {
							stack.Push(tmp)
							break
						}
					}
				}
			}

		} else {
			datas = append(datas, tmp)
		}
	}

	//将栈内剩余的操作符全部弹出。
	for {
		if pop := stack.Pop(); pop != nil {
			datas = append(datas, pop.(string))
		} else {
			break
		}
	}
	return datas
}

//获取运算符优先级
func getOperatorLevel(str string) (level int) {
	switch str {
	case "+", "-":
		level = 0
	case "*", "/":
		level = 1
	case "^":
		level = 2
	}
	return level
}

//比较优先级
// if return 1, o1 > o2.
// if return 0, o1 = 02
// if return -1, o1 < o2
func compareOperator(o1, o2 string) int {

	o1Priority := getOperatorLevel(o1)
	o2Priority := getOperatorLevel(o2)

	if o1Priority > o2Priority {
		return 1
	} else if o1Priority == o2Priority {
		return 0
	} else {
		return -1
	}
}

func IsNumberString(o1 string) bool {
	if o1 == "+" || o1 == "-" || o1 == "*" || o1 == "/" || o1 == "(" || o1 == ")" || o1 == "^" {
		return false
	} else {
		return true
	}
}

func isNumber(o1 byte) bool {
	if o1 == '+' || o1 == '-' || o1 == '*' || o1 == '/' || o1 == '(' || o1 == ')' || o1 == '^' {
		return false
	} else {
		return true
	}
}

//将字符串转换成字符串数组
func convertToStrings(s string) []string {
	var strs []string
	bys := []byte(s)
	var tmp string
	for i := 0; i < len(bys); i++ {
		if !isNumber(bys[i]) {
			if tmp != "" {
				strs = append(strs, tmp)
				tmp = ""
			}
			strs = append(strs, string(bys[i]))
		} else {
			tmp = tmp + string(bys[i])
		}
	}
	strs = append(strs, tmp)
	return strs
}
