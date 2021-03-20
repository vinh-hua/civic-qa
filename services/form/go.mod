module github.com/team-ravl/civic-qa/services/form

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/team-ravl/civic-qa/services/common v0.0.0
	github.com/team-ravl/civic-qa/services/logAggregator v0.0.0-20210301202124-142c59451b9f
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)

replace (
	github.com/team-ravl/civic-qa/services/common v0.0.0 => ../common
	github.com/team-ravl/civic-qa/services/logAggregator v0.0.0-20210301202124-142c59451b9f => ../logAggregator
)
