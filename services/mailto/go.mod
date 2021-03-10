module github.com/vivian-hua/civic-qa/services/mailto

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0
)

replace (
	github.com/vivian-hua/civic-qa/services/common v0.0.0 => ../common
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0 => ../logAggregator
)
