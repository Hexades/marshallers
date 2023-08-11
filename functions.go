package marshallers

import (
	"encoding/json"
	"log"
)

type marshal func(model any) []byte
type unmarshal func(model any, bytes []byte)

var marshalJSON = func(model any) []byte {
	b, _ := json.Marshal(model)
	return b
}

var unmarhsalJSON = func(bytes []byte, model any) {
	log.Println(bytes)
	json.Unmarshal(bytes, model)
}
