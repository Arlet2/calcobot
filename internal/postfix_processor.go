package internal

import (
	"errors"
	"regexp"
	"strings"
)

type postfixExpression string

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
		if '0' <= value && value <= '9' {
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

	return postfixExpression(output), nil
}