package event

import (
	"github.com/Ontology/common"
)

type LogEventArgs struct {
	Container common.Uint256
	CodeHash  common.Address
	Message   string
}