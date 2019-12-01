package types

type TxResponseResult struct {
	Height    int64           `json:"height"`
	TxHash    string          `json:"txhash"`
	Code      uint32          `json:"code,omitempty"`
	Data      string          `json:"data,omitempty"`
	RawLog    string          `json:"raw_log,omitempty"`
	Info      string          `json:"info,omitempty"`
	GasWanted int64           `json:"gas_wanted,omitempty"`
	GasUsed   int64           `json:"gas_used,omitempty"`
	Codespace string          `json:"codespace,omitempty"`
	Tx        string              `json:"tx,omitempty"`
	Timestamp string          `json:"timestamp,omitempty"`
}
