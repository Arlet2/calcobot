package model

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type postfixExpression string

// Convert user input string to postfix expression
// Can produce error "Operation is not allowed" if operation doesn't support for this converter
func ToPostfix(input string) (postfixExpression, error) {
	var stack Stack[operation] = &OperationStack{} 
	output := ""
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
				Если операция допускается, то выталкиваем все операции с бОльшим либо равным приоритетом
			*/
			operation, err := AllowedOperations.GetOperation(value)

			if err != nil {
				return "", err
			}

			addedOperation := operation

			if !stack.IsEmpty() && value != '(' {
				operation, err = stack.Peek()

				for operation.priority >= addedOperation.priority {
					if err != nil {
						break
					}
					stack.Pop()
					output += string(operation.symbol) + " "
					operation, err = stack.Peek()
				}
			}

			stack.Push(addedOperation)
		} else {
			return "", errors.New(string(value)+" operation is not allowed")
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

	output = strings.TrimSpace(output)

	return postfixExpression(output), nil
}

// Calculate postfix expression and return result of the calculation
// Can produce error "Operation is not allowed" if operation is not implemented in dictionary of allowed operations
func CalculatePostfix(exp postfixExpression) (float64, error) {
	var stack Stack[float64] = &NumberStack{}

	input := string(exp)
	number := ""

	for _, value := range input {

		if '0' <= value && value <= '9' || value == '.' {
			number += string(value)
			continue
		} else if number != "" {
			if len(number) > 307 {
				return 0, errors.New("одно из чисел слишком большое")
			}
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
