package operator

import (
	"fmt"

	"github.com/rodrigobotti/go-lessons/packages/prefix"
)

var (
	// OperatorName is the operator's name
	OperatorName = "NXTelCo"
	// OperatorPrefix just a value
	OperatorPrefix = fmt.Sprintf("%d %s", prefix.Capitol, OperatorName)
)
