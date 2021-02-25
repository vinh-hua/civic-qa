module github.com/vivian-hua/civic-qa/services/gateway

go 1.15

require (
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0-20210224044125-3ddd4a4c5df4
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0
)

replace (
	github.com/vivian-hua/civic-qa/services/common v0.0.0-20210213001141-bd4fbd60d179 => ../common
	github.com/vivian-hua/civic-qa/services/logAggregator v0.0.0 => ../logAggregator
)
