// Copyright 2018 The QOS Authors

package cmd

import (
	"github.com/QOSGroup/qmoon_qos_agent/module/block"
	"github.com/QOSGroup/qmoon_qos_agent/module/distribution"
	"github.com/QOSGroup/qmoon_qos_agent/module/gov"
	"github.com/QOSGroup/qmoon_qos_agent/module/mint"
	"github.com/QOSGroup/qmoon_qos_agent/module/stake"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// ServerCmd qmoon http server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "http server",
	RunE:  server,
}

var (
	laddr string
)

func init() {
	ServerCmd.PersistentFlags().StringVar(&laddr, "laddr", "0.0.0.0:19528", "listen addr")
}

type TxQuery struct {
	Txs []string `json:"txs"`
}

func server(cmd *cobra.Command, args []string) error {
	r := gin.Default()
	block.Register(r)
	gov.Register(r)
	distribution.Register(r)
	mint.Register(r)
	stake.Register(r)
	return r.Run(laddr)
}
