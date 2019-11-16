package stdlib

import (
	"encoding/hex"
	"github.com/n-is/tengo/objects"
)

var hexModule = map[string]objects.Object{
	"encode": &objects.UserFunction{Value: FuncAYRS(hex.EncodeToString)},
	"decode": &objects.UserFunction{Value: FuncASRYE(hex.DecodeString)},
}
