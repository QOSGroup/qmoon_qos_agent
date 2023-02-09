module github.com/QOSGroup/qmoon_qos_agent

go 1.12

require (
	github.com/QOSGroup/qbase v0.2.3-0.20190927065041-32eb90018d34
	github.com/QOSGroup/qos v0.0.8
	github.com/gin-gonic/gin v1.7.7
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.3.2
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/tendermint v0.32.2
)

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190801041406-cbf593c0f2f3
