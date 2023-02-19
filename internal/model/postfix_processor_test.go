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
		"3 !", // проверка факториала
		"5 ! 4 ! /", // проверка унарности факториала
		"4 2 2 * 2 / + 2 - 1 ~ +", // проверка одинаковый приоритетов
	}

	answers := []float64{
		130,
		2,
		-2,
		0,
		6,
		5,
		3,
	}

	t.Log("Start testing calculating postfix expressions.")
	{
		for testID, value := range testExpressions {
			t.Logf("\tTest %d:\tcalculating "+string(value), testID)
			answer, err := CalculatePostfix(postfixExpression(value))
			
			if err != nil {
				t.Logf("Found error in test %d: "+err.Error(), testID)
				t.Fail()
			}

			if answer != answers[testID] {
				t.Logf("Wrong answer in test %d: expected %f, found: %f", testID, answers[testID], answer)
				t.Fail()
			}
		}

	}
}

func TestParsing(t *testing.T) {
	testExpressions := []string{
		"2+2",
		"4+2*4",
		"4+2*2/2-2+~1",
		"2*(3+2)",
		"2^(2+4)",
	}

	answers := []string{
		"2 2 +",
		"4 2 4 * +",
		"4 2 2 * 2 / + 2 - 1 ~ +",
		"2 3 2 + *",
		"2 2 4 + ^",
	}

	t.Log("Start testing parsing expressions.")
	{
		for testID, value := range testExpressions {
			t.Logf("\tTest %d:\tparsing "+string(value), testID)
			answer, err := ToPostfix(value)
			
			if err != nil {
				t.Logf("Found error in test %d: "+err.Error(), testID)
				t.Fail()
			}

			if string(answer) != answers[testID] {
				t.Logf("Wrong answer in test %d: expected %s, found: %s", testID, answers[testID], answer)
				t.Fail()
			}
		}

	}
}

func TestAllProcess(t *testing.T) {

	testExpressions := []string{
		"2+2",
		"4+2*4",
		"2*(3+2)",
		"2^(2+4)",
		"~1^2",
	}

	answers := []float64{
		4,
		12,
		10,
		64,
		1,
	}

	t.Log("Start testing parsing and calculating expressions.")
	{
		for testID, value := range testExpressions {
			t.Logf("\tTest %d:\tparsing "+string(value), testID)
			postfixExpression, err := ToPostfix(value)
			
			if err != nil {
				t.Logf("Found error in test %d: "+err.Error(), testID)
				t.Fail()
			}

			answer, err := CalculatePostfix(postfixExpression)

			if err != nil {
				t.Logf("Found error in test %d: "+err.Error(), testID)
				t.Fail()
			}

			if answer != answers[testID] {
				t.Logf("Wrong answer in test %d: expected %f, found: %f", testID, answers[testID], answer)
				t.Fail()
			}
		}

	}
}