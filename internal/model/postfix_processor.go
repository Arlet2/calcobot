package model

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type postfixExpression string

//todo починить!
func ToPostfix(input string) (postfixExpression, error) {
	output := ""
	stack := OperationStackImpl{}
	number := ""

	// очищаем от пробелов
	input = strings.ReplaceAll(input, " ", "")

	for _, value := range input {

		/*
			Конструируем число из цифр
		*/
		if ('0' <= value && value <= '9') || value == '.' || value == ',' {
			number += string(value)
			continue
		} else {
			output += number + " "
			number = ""
		}

		/*
			При закрывающейся скобке выталкиваем все операции до открывающей скобки
			Если там была закрывающая скобка, то значит скобки поставлены неверно
		*/

		if value == ')' {
			operation, err := stack.Pop()

			for operation.symbol != '(' {
				if err != nil {
					break
				}
				if operation.symbol == ')' {
					return "", errors.New("incorrect bracket position")
				}
				output += string(operation.symbol) + " "
				operation, err = stack.Pop()
			}

		} else if AllowedOperations.IsOnDict(value) {
			/*
				Если операция допускается, то выталкиваем все операции с бОльшим приоритетом
			*/
			operation, err := AllowedOperations.GetOperation(value)

			if err != nil {
				return "", err
			}

			addedOperation := operation

			if !stack.IsEmpty() {
				operation, err = stack.Peek()

				for operation.priority > addedOperation.priority {
					if err != nil {
						break
					}
					stack.Pop()
					output += string(operation.symbol) + " "
					operation, err = stack.Peek()
				}
			}

			stack.Push(addedOperation)
		}
	}
	output += number + " "

	operations, err := stack.PopAll()

	if err == nil {
		for _, value := range operations {
			output += string(value.symbol) + " "
		}
	}

	// очищаем от избыточных пробелов
	var regex = regexp.MustCompile(`\s{2,}`)
	output = regex.ReplaceAllString(output, " ")

	// преобразуем , в .
	output = strings.ReplaceAll(output, ",", ".")

	return postfixExpression(output), nil
}

func CalculatePostfix(exp postfixExpression) (float64, error) {
	input := string(exp)
	stack := NumberStackImpl{}
	number := ""

	for _, value := range input {

		if '0' <= value && value <= '9' || value == '.' {
			number += string(value)
			continue
		} else if number != "" {
			parsedNumber, _ := strconv.ParseFloat(number, 64)
			number = ""
			stack.Push(parsedNumber)
		}

		if value == ' ' {
			continue
		}

		operation, err := AllowedOperations.GetOperation(value)

		if err != nil {
			return 0, err
		} else {
			b, err := stack.Pop()

			if err != nil {
				return 0, err
			}

			if operation.isBinary {

				a, err := stack.Pop()

				if err != nil {
					return 0, err
				}

				stack.Push(operation.calc(a, b))
			} else {
				stack.Push(operation.calc(b, 0))
			}
		}
	}

	result, err := stack.Pop()

	return result, err
}
