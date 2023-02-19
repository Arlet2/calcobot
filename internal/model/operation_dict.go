package model

import "errors"

type OperationDict interface {
	IsOnDict(symbol rune) bool
	GetPriority(symbol rune) (int, error)
	GetOperation(symbol rune) (operation, error)
}

type OperationDictImpl struct{
	operations []operation
}

func (dict OperationDictImpl) IsOnDict(symbol rune) bool {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return true
		}
	}
	return false
}

func (dict OperationDictImpl) GetPriority(symbol rune) (int, error) {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return value.priority, nil
		}
	}

	return 0, errors.New(string(symbol)+" operation is not allowed")
}

func (dict OperationDictImpl) GetOperation(symbol rune) (operation, error) {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return value, nil
		}
	}

	return operation{}, errors.New(string(symbol)+" operation is not allowed")
}