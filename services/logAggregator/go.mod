module github.com/vivian-hua/civic-qa/services/logAggregator

go 1.15

require (
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v1.14.6 // indirect
	github.com/urfave/negroni v1.0.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0-20210301202124-142c59451b9f
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.12
)

replace github.com/vivian-hua/civic-qa/services/common v0.0.0-20210227211936-aef745c47c5f => ../common
