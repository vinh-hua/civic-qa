module github.com/team-ravl/civic-qa/services/gateway

go 1.15

require (
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/team-ravl/civic-qa/services/common v0.0.0
	github.com/team-ravl/civic-qa/services/logAggregator v0.0.0
)

replace (
	github.com/team-ravl/civic-qa/services/common v0.0.0 => ../common
	github.com/team-ravl/civic-qa/services/logAggregator v0.0.0 => ../logAggregator
)
