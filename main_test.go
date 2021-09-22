package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Transactions []*Transaction
	Result       []*Transaction
	Interval     string
	Err          error
}

func TestFormatGraphs(t *testing.T) {
	transactions := []*Transaction{
		{
			4456,
			1616026248,
		},
		{
			4231,
			1616022648,
		},
		{
			5212,
			1616019048,
		},
		{
			4321,
			1615889448,
		},
		{
			4567,
			1615871448,
		},
	}

	expectedResult := []*Transaction{
		{
			4456,
			1616025600,
		},
		{
			4231,
			1615939200,
		},
		{
			4321,
			1615852800,
		},
	}
	test := &testCase{
		Transactions: transactions,
		Result:       expectedResult,
		Interval:     "day",
		Err:          nil,
	}

	res, err := formatGraphs(test.Transactions, test.Interval)

	if err != test.Err {
		t.Errorf("Expected error =  %v, got error =  %v", test.Err, err)
	}

	assert.ElementsMatch(t, test.Result, res)

}
func TestFormatGraphsErrorUnknownInterval(t *testing.T) {
	transactions := []*Transaction{
		{
			4456,
			1616026248,
		},
		{
			4231,
			1616022648,
		},
		{
			5212,
			1616019048,
		},
		{
			4321,
			1615889448,
		},
		{
			4567,
			1615871448,
		},
	}

	expectedResult := []*Transaction{
		{
			4456,
			1616025600,
		},
		{
			4231,
			1615939200,
		},
		{
			4321,
			1615852800,
		},
	}
	test := &testCase{
		Transactions: transactions,
		Result:       expectedResult,
		Interval:     "unknown_interval",
		Err:          errUnknownInterval,
	}

	_, err := formatGraphs(test.Transactions, test.Interval)
	if err != test.Err {
		t.Errorf("Expected error =  %v, got error =  %v", test.Err, err)
	}

}

func TestFormatGraphsErrorEmptyTransactionInput(t *testing.T) {
	test := &testCase{
		Transactions: nil,
		Result:       nil,
		Interval:     "day",
		Err:          errEmptyTransactionInput,
	}

	_, err := formatGraphs(test.Transactions, test.Interval)
	if err != test.Err {
		t.Errorf("Expected error =  %v, got error =  %v", test.Err, err)
	}
}

func TestFormatGraphsErrorEmptyIntervalInput(t *testing.T) {
	transactions := []*Transaction{
		{
			4456,
			1616026248,
		},
		{
			4231,
			1616022648,
		},
		{
			5212,
			1616019048,
		},
		{
			4321,
			1615889448,
		},
		{
			4567,
			1615871448,
		},
	}
	test := &testCase{
		Transactions: transactions,
		Result:       nil,
		Interval:     "",
		Err:          errEmptyIntervalInput,
	}

	_, err := formatGraphs(test.Transactions, test.Interval)
	if err != test.Err {
		t.Errorf("Expected error =  %v, got error =  %v", test.Err, err)
	}
}
