module moriaty.com/cia/cia-executor

go 1.13

require (
	github.com/go-redis/redis v6.15.7+incompatible
	github.com/micro/go-micro v1.18.0
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	moriaty.com/cia/cia-common v0.0.0-incompatible
)

replace moriaty.com/cia/cia-common => ../CIA-Common
