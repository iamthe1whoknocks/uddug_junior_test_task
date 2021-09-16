package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	errUnknownInterval = errors.New("Unknown interval provided, allowed intervals are : month, week, day, hour")
	errEmptyInput      = errors.New("Empty input provided")
)

// Transaction is user transaction in exchange
type Transaction struct {
	Value     int
	Timestamp int64 //in example you provide this field is not time.Time struct
}

func main() {
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
	formattedGraphs, err := formatGraphs(transactions, "day")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, s := range formattedGraphs {
		fmt.Printf("%+v\n", s)
	}
}

// formatGraph function groups structures by timestamp (rounding it up to the required time interval)
// and returns last value in this time interval
func formatGraphs(transactions []*Transaction, intervalStr string) (
	[]*Transaction,
	error) {
	if transactions == nil {
		return nil, errEmptyInput
	}
	interval, err := translateInterval(strings.ToLower(intervalStr))
	if err != nil {
		return nil, err
	}
	m := make(map[time.Time]int)
	for _, t := range transactions {
		v, ok := m[time.Unix(t.Timestamp, 0).Truncate(interval)]
		if !ok {
			m[time.Unix(t.Timestamp, 0).Truncate(interval)] = t.Value
		}
		if t.Value < v {
			m[time.Unix(t.Timestamp, 0).Truncate(interval)] = t.Value
		}
	}
	formattedGraphs := make([]*Transaction, 0, len(m))
	for k, v := range m {
		transaction := &Transaction{
			Value:     v,
			Timestamp: k.Unix(),
		}
		formattedGraphs = append(formattedGraphs, transaction)

	}
	return formattedGraphs, nil
}

func translateInterval(intervalStr string) (interval time.Duration, err error) {
	switch intervalStr {
	case "month":
		return time.Hour * 24 * 31, nil
	case "week":
		return time.Hour * 24 * 7, nil
	case "day":
		return time.Hour * 24, nil
	case "hour":
		return time.Hour, nil
	default:
		return 0, errUnknownInterval
	}
}
