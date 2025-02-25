module github.com/miseyu/gen/tests

go 1.16

require (
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	golang.org/x/tools v0.5.0 // indirect
	gorm.io/driver/mysql v1.4.5
	gorm.io/driver/sqlite v1.4.4
	github.com/miseyu/gen v0.3.19
	gorm.io/gorm v1.24.3
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.4.0
)

replace github.com/miseyu/gen => ../
