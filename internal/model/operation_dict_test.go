package model

import "testing"

func TestDict(t *testing.T) {
	dict := OperationDictImpl{
		[]operation{
			{'+', 2, true, add},
			{'-', 2, true, sub},
		},
	}

	t.Log("Start testing operation of OperationDict.")
	{
		testID := 0

		t.Logf("\tTest %d: check existing operations", testID)
		{
			if !dict.IsOnDict('+') {
				t.Logf("Fail on test %d: + operation is not found at dict", testID)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: check not existing operations", testID)
		{
			if dict.IsOnDict('/') {
				t.Logf("Fail on test %d: / operation is found on dict but does not exist", testID)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: get operations by symbol", testID)
		{
			operation, err := dict.GetOperation('+')

			if err != nil {
				t.Logf("Fail on test %d: found error: "+err.Error(), testID)
				t.Fail()
			}

			if operation.symbol != '+' {
				t.Logf("Fail on test %d: expected: + but found: %c", testID, operation.symbol)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: get operations by not existing symbol", testID)
		{
			_, err := dict.GetOperation('/')

			if err == nil {
				t.Logf("Fail on test %d: expected error", testID)
				t.Fail()
			}
		}
		testID++
		
		t.Logf("\tTest %d: get priority by symbol", testID)
		{
			priority, err := dict.GetPriority('+')

			if err != nil {
				t.Logf("Fail on test %d: found error: "+err.Error(), testID)
				t.Fail()
			}

			if priority != 2 {
				t.Logf("Fail on test %d: expected 2 but found %d", testID, priority)
				t.Fail()
			}
		}
		testID++

		t.Logf("\tTest %d: get priority by not existing symbol", testID)
		{
			_, err := dict.GetPriority('/')

			if err == nil {
				t.Logf("Fail on test %d: expected error", testID)
				t.Fail()
			}
		}
		testID++
	}
}