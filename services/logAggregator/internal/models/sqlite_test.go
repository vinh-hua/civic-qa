package models

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	// random test UUIDs
	testUUID1 = uuid.MustParse("7ee5d007-a780-477c-988f-32faf595045f")
	testUUID2 = uuid.MustParse("6aa4d006-a670-466c-877f-21faf484034f")
)

// Tests logging
func TestLog(t *testing.T) {
	cases := []LogEntry{
		{CorrelationID: testUUID1},
		{CorrelationID: testUUID1, Service: "test"},
		{CorrelationID: testUUID1, Service: "test", Notes: "test"},
		{CorrelationID: testUUID1, TimeUnix: time.Now().Unix(), Service: "test", Notes: "test"},
		{CorrelationID: testUUID1, TimeUnix: time.Now().Unix(), Service: "test", StatusCode: 200, Notes: "test"},
	}

	store := NewSQLiteLogStore(":memory:", true)
	for i, entry := range cases {
		err := store.Log(entry)
		if err != nil {
			t.Fatalf("Error logging %d: %v", i, err)
		}
	}
}

// Tests query
func TestQuery(t *testing.T) {
	type casesStruct struct {
		data  []LogEntry
		query LogQuery
		out   []LogEntry
	}

	cases := []casesStruct{
		// 0 single log single out
		{

			data: []LogEntry{
				{CorrelationID: testUUID1, Service: "test"},
			},
			query: LogQuery{testUUID1, 0, 0, "", 0, 0},
			out: []LogEntry{
				{CorrelationID: testUUID1, Service: "test"},
			},
		},
		// 1 two log two out
		{
			data: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
			query: LogQuery{testUUID1, 0, 0, "", 0, 0},
			out: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
		},
		// 2 two log one out (service)
		{
			data: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
			query: LogQuery{testUUID1, 0, 0, "A", 0, 0},
			out: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
			},
		},
		// 3 two log one out (status)
		{
			data: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1, StatusCode: 500},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
			query: LogQuery{testUUID1, 0, 0, "A", 400, 0},
			out: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1, StatusCode: 500},
			},
		},
		// 4 two log one out (uuid)
		{
			data: []LogEntry{
				{CorrelationID: testUUID1, Service: "A"},
				{CorrelationID: testUUID2, Service: "B"},
			},
			query: LogQuery{testUUID1, 0, 0, "", 0, 0},
			out: []LogEntry{
				{CorrelationID: testUUID1, Service: "A"},
			},
		},
		// 5 two log two out (no uuid)
		{
			data: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID2, Service: "B", TimeUnix: 2},
			},
			query: LogQuery{},
			out: []LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID2, Service: "B", TimeUnix: 2},
			},
		},
	}

	for i, testCase := range cases {
		// log initial data
		store := NewSQLiteLogStore(":memory:", true)
		for _, entry := range testCase.data {
			err := store.Log(entry)
			if err != nil {
				t.Fatalf("Error logging %d: %v", i, err)
			}
		}

		// query
		result, err := store.Query(testCase.query)
		if err != nil {
			t.Fatalf("Error querying %d: %v", i, err)
		}

		// check matches (order matters)
		for i := range result {
			if result[i] != testCase.out[i] {
				t.Fatalf("Unmatched result: got \n\t%v expected \n\t%v", result[i], testCase.out[i])
			}
		}

		// ensure length matches
		if len(result) != len(testCase.out) {
			t.Fatalf("Unexpected result size: expected %d got %d", len(result), len(testCase.out))
		}
	}
}

// Tests SQL generator cases
func TestSQLGenerator(t *testing.T) {
	cases := []LogQuery{
		{},
		{testUUID1, 0, 0, "", 0, 0},
		{uuid.Nil, 1, 10, "", 0, 0},
		{uuid.Nil, 0, 0, "A", 0, 0},
		{uuid.Nil, 0, 0, "", 100, 500},
		{testUUID1, time.Now().Unix(), 0, "test", 0, 399},
		{testUUID1, time.Now().Unix(), time.Now().Unix() + 5, "test", 1, 399},
	}

	store := NewSQLiteLogStore(":memory:", true)

	for i, query := range cases {
		_, _, err := store.generateSQL(query)
		if err != nil {
			t.Fatalf("Error generating %d: %v", i, err)
		}
	}
}
