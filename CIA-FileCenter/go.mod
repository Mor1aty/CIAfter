module moriaty.com/cia/cia-filecenter

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/satori/go.uuid v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	moriaty.com/cia/cia-common v0.0.0-incompatible
)

replace moriaty.com/cia/cia-common => ../CIA-Common
