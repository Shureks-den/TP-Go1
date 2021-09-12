package calculator

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

type customStack struct {
	stack *list.List
}

func (c *customStack) Push(value string) {
	c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *customStack) Front() (string, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("Peep Error: Stack Datatype is incorrect")
	}
	return "", fmt.Errorf("Peep Error: Stack is empty")
}

func (c *customStack) Size() int {
	return c.stack.Len()
}

func (c *customStack) Empty() bool {
	return c.stack.Len() == 0
}

func GetPreparedData() []string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str := sc.Text()
	str = TidyString(str)
	c := strings.Split(str, "")
	return c
}

func TidyString(str string) string {
	runeValue, width := utf8.DecodeRuneInString(str[:])
	if runeValue == '"' {
		str = str[width:]
	}
	runeValue, width = utf8.DecodeRuneInString(str[len(str)-1:])
	if runeValue == '"' {
		str = str[:len(str)-1]
	}
	r := regexp.MustCompile("[\\sa-zA-Z]+")
	str = r.ReplaceAllString(str, "")
	return str
}

func isDigit(sign string) bool {
	_, er := strconv.Atoi(sign)
	if er == nil {
		return true
	}
	return false
}

func getPriority(sign string) int {
	switch sign {
	case "+":
		return 1
	case "-":
		return 1
	case "*":
		return 2
	case "/":
		return 2
	default:
		return 0
	}
}

func pushNum(varStack *customStack, inRow int, value string) {
	var sb strings.Builder
	if s, _ := varStack.Front(); inRow != 0 {
		sb.WriteString(s)
		varStack.Pop()
	}
	sb.WriteString(value)
	varStack.Push(sb.String())
}

func calculation(varStack *customStack, funcStack *customStack) error {
	elem, err := funcStack.Front()
	if err != nil {
		return err
	}
	funcStack.Pop()

	firstNum, err := varStack.Front()
	if err != nil {
		return err
	}
	varStack.Pop()

	secondNum, err := varStack.Front()
	varStack.Pop()
	if err != nil {
		return err
	}

	val1, _ := strconv.Atoi(firstNum)
	val2, _ := strconv.Atoi(secondNum)
	var res int
	switch elem {
	case "+":
		res = val2 + val1
	case "-":
		res = val2 - val1
	case "/":
		res = val2 / val1
	case "*":
		res = val2 * val1
	}
	varStack.Push(strconv.Itoa(res))
	return nil
}

func Calculator(data []string) (string, error) {
	varStack := &customStack{
		stack: list.New(),
	}
	funcStack := &customStack{
		stack: list.New(),
	}

	var inRow int
	for _, value := range data {
		if i := isDigit(value); i {
			// добавляем число
			pushNum(varStack, inRow, value)
			inRow++
		} else {
			// вычисление выражения
			if value != "(" && value != ")" {
				if fr, err := funcStack.Front(); getPriority(fr) >= getPriority(value) {
					if err != nil {
						return "", err
					}
					calculation(varStack, funcStack)
				}
			}
			// если нашли закрывающую скобку вычисляем до открывающей
			if value == ")" {
				a, _ := funcStack.Front()
				for a != "(" {
					calculation(varStack, funcStack)
					a, _ = funcStack.Front()
				}
				funcStack.Pop()
			} else {
				// если это открывающая скобка или арифметическое действие
				funcStack.Push(value)
			}
			// в любом случае обнуляем счетчик цифр
			inRow = 0
		}
	}

	for !funcStack.Empty() {
		calculation(varStack, funcStack)
	}
	return varStack.Front()
}
