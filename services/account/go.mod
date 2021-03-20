module github.com/team-ravl/civic-qa/service/account

go 1.15

require (
	github.com/go-redis/redis v6.15.9+incompatible // indirect
	github.com/go-redis/redis/v8 v8.7.0
	github.com/gorilla/mux v1.8.0
	github.com/team-ravl/civic-qa/services/common v0.0.0
	github.com/team-ravl/civic-qa/services/logAggregator v0.0.0-20210227211936-aef745c47c5f
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)

replace (
	github.com/team-ravl/civic-qa/services/common v0.0.0 => ../common
	github.com/team-ravl/civic-qa/services/logAggregator v0.0.0-20210227211936-aef745c47c5f => ../logAggregator
)
