package model

import (
	"testing"
)

func TestCalculating(t *testing.T) {

	testExpressions := []postfixExpression{
		"2 2 4 + ^ 2 * 2 +", // проверка приоритетов
		"4 2 /", // проверка корректности деления
		"4 ~ 2 +", // проверка унарного оператора
		"0 4 2 - *", // проверка умножения на 0
	}

	answers := []float64{
		130,
		2,
		-2,
		0,
	}

	t.Log("Start testing calculating of postfix expressions.")
	{
		for testID, value := range testExpressions {
			t.Logf("\tTest %d:\tcalculating "+string(testExpressions[testID]), testID)
			answer, err := CalculatePostfix(postfixExpression(value))
			
			if err != nil {
				t.Log("Found error in test %d: "+err.Error(), testID)
				t.Fail()
			}

			if answer != answers[testID] {
				t.Logf("Wrong answer in test %d: expected %f, found: %f", testID, answers[testID], answer)
				t.Fail()
			}
		}

	}
}
