module github.com/vivian-hua/civic-qa/service/account

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/vivian-hua/civic-qa/services/common v0.0.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.12
)

replace github.com/vivian-hua/civic-qa/services/common v0.0.0 => ../common
