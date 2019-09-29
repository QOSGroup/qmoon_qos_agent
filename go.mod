module github.com/QOSGroup/qmoon_qos_agent

go 1.12

require (
	github.com/QOSGroup/qbase v0.2.3-0.20190923023519-41e227af6e4c
	github.com/QOSGroup/qos v0.0.6-0.20190926011623-143f93e8a9d1
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.2
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190801041406-cbf593c0f2f3
