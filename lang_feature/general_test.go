package lang_feature

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPassingByCopy(t *testing.T) {
	amount := 6

	applyCopy(amount)

	assert.Equal(t, amount, 6)
}

func TestPassingByRef(t *testing.T) {
	amount := 6

	applyRef(&amount)

	assert.NotEqual(t, amount, 6)
	assert.Equal(t, amount, 12)
}

func TestPtrsGeneral01(t *testing.T) {
	var myint int
	assert.Equal(t, "*int", reflect.TypeOf(&myint).String())

	var myfloat32 float32
	assert.Equal(t, "*float32", reflect.TypeOf(&myfloat32).String())

	var mybool bool
	assert.Equal(t, "*bool", reflect.TypeOf(&mybool).String())

	var device = map[string]string{"one": "1", "two": "2"}
	assert.Equal(t, "*map[string]string", reflect.TypeOf(&device).String())

	checkMapPassingByRef(&device)
	assert.Equal(t, device["one"], "2")
}
