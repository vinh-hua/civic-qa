module github.com/vivian-hua/civic-qa/services/auth

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0-20210217005008-848fa3bc0399
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0-20210223013533-594b6884c79e
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.12
)

replace (
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0-20210223013533-594b6884c79e => ../logAggregator

)
