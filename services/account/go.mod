module github.com/vivian-hua/civic-qa/service/account

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0
)

replace github.com/vivian-hua/civic-qa/services/common v0.0.0 => ../common
