package repository

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/vivian-hua/civic-qa/services/logAggregator/internal/model"
)

var (
	// random test UUIDs
	testUUID1 = uuid.MustParse("7ee5d007-a780-477c-988f-32faf595045f")
	testUUID2 = uuid.MustParse("6aa4d006-a670-466c-877f-21faf484034f")
	lgr       = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:      logger.Silent, //
			Colorful:      true,          //
			SlowThreshold: 0 * time.Microsecond,
		},
	)
)

func logEqual(log1, log2 model.LogEntry) bool {

	return log1.CorrelationID == log2.CorrelationID &&
		log1.TimeUnix == log2.TimeUnix &&
		log1.Service == log2.Service &&
		log1.StatusCode == log2.StatusCode &&
		log1.Notes == log2.Notes
}

func createRepo() *LogRepository {
	// Handler context
	repo, err := NewLogRepository(sqlite.Open(":memory:"), &gorm.Config{Logger: lgr})
	if err != nil {
		panic(err)
	}
	return repo
}

// Tests logging
func TestLog(t *testing.T) {
	cases := []model.LogEntry{
		{CorrelationID: testUUID1},
		{CorrelationID: testUUID1, Service: "test"},
		{CorrelationID: testUUID1, Service: "test", Notes: "test"},
		{CorrelationID: testUUID1, TimeUnix: time.Now().Unix(), Service: "test", Notes: "test"},
		{CorrelationID: testUUID1, TimeUnix: time.Now().Unix(), Service: "test", StatusCode: 200, Notes: "test"},
	}

	repo := createRepo()
	for i, entry := range cases {
		err := repo.Log(entry)
		if err != nil {
			t.Fatalf("Error logging %d: %v", i, err)
		}
	}
}

// Tests query
func TestQuery(t *testing.T) {
	type casesStruct struct {
		data  []model.LogEntry
		query model.LogQuery
		out   []model.LogEntry
	}

	cases := []casesStruct{
		// 0 single log single out
		{
			data: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "test"},
			},
			query: model.LogQuery{testUUID1, 0, 0, "", 0, 0},
			out: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "test"},
			},
		},
		// 1 two log two out
		{
			data: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
			query: model.LogQuery{testUUID1, 0, 0, "", 0, 0},
			out: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
		},
		// 2 two log one out (service)
		{
			data: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
			query: model.LogQuery{testUUID1, 0, 0, "A", 0, 0},
			out: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
			},
		},
		// 3 two log one out (status)
		{
			data: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1, StatusCode: 500},
				{CorrelationID: testUUID1, Service: "B", TimeUnix: 2},
			},
			query: model.LogQuery{testUUID1, 0, 0, "A", 400, 0},
			out: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1, StatusCode: 500},
			},
		},
		// 4 two log one out (uuid)
		{
			data: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A"},
				{CorrelationID: testUUID2, Service: "B"},
			},
			query: model.LogQuery{testUUID1, 0, 0, "", 0, 0},
			out: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A"},
			},
		},
		// 5 two log two out (no uuid)
		{
			data: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID2, Service: "B", TimeUnix: 2},
			},
			query: model.LogQuery{},
			out: []model.LogEntry{
				{CorrelationID: testUUID1, Service: "A", TimeUnix: 1},
				{CorrelationID: testUUID2, Service: "B", TimeUnix: 2},
			},
		},
	}

	for i, testCase := range cases {
		// log initial data
		repo := createRepo()
		for _, entry := range testCase.data {
			err := repo.Log(entry)
			if err != nil {
				t.Fatalf("Error logging %d: %v", i, err)
			}
		}

		// query
		result, err := repo.Query(testCase.query)
		if err != nil {
			t.Fatalf("Error querying %d: %v", i, err)
		}

		// check matches (order matters)
		for i := range result {
			if !logEqual(result[i], testCase.out[i]) {
				t.Fatalf("Unmatched result: got \n\t%v expected \n\t%v", result[i], testCase.out[i])
			}
		}

		// ensure length matches
		if len(result) != len(testCase.out) {
			t.Fatalf("case %d: Unexpected result size: expected %d got %d", i, len(testCase.out), len(result))
		}
	}
}
