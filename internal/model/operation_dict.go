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

// Check existing of operation in allowed operations.
// Return true if operation exists in dictionary, otherwise false.
func (dict OperationDictImpl) IsOnDict(symbol rune) bool {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return true
		}
	}
	return false
}

// Get priority of some operation by symbol (char).
// Can produce error "Operation is not allowed" if this operation does not implemented.
func (dict OperationDictImpl) GetPriority(symbol rune) (int, error) {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return value.priority, nil
		}
	}

	return 0, errors.New(string(symbol)+" operation is not allowed")
}

// Get operation object by symbol (char).
// Can produce error "Operation is not allowed" if this operation does not implemented.
func (dict OperationDictImpl) GetOperation(symbol rune) (operation, error) {
	for _, value := range dict.operations {
		if value.symbol == symbol {
			return value, nil
		}
	}

	return operation{}, errors.New(string(symbol)+" operation is not allowed")
}