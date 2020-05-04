module moriaty.com/cia/cia-executor

go 1.13

require (
	github.com/micro/go-micro v1.18.0
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	gopkg.in/ini.v1 v1.44.0
	moriaty.com/cia/cia-common v0.0.0-incompatible
)

replace moriaty.com/cia/cia-common => ../CIA-Common
