module github.com/vivian-hua/civic-qa/service/account

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0-20210301202124-142c59451b9f
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0-20210227211936-aef745c47c5f
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.12
)

replace (
	github.com/vivian-hua/civic-qa/services/common v0.0.0-20210227211936-aef745c47c5f => ../common
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0-20210227211936-aef745c47c5f => ../logAggregator
)
