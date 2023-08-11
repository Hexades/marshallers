package marshallers

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type TestData struct {
	Name  string `json:"name"`
	Years int    `json:"int"`
}

func TestJSONMarshaling(t *testing.T) {
	td := TestData{"Brad", 25}
	bytes := marshalJSON(td)
	assert.NotNil(t, bytes)
	model := TestData{}
	unmarhsalJSON(bytes, &model)
	log.Println(model)
	assert.Equal(t, td, model)
}
