module moriaty.com/cia/cia-supporter

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/micro/go-micro v1.18.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.52.0
	moriaty.com/cia/cia-common v0.0.0-incompatible
)

replace moriaty.com/cia/cia-common => ../CIA-Common
