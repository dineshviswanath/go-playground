package kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertCamelCaseDash(t *testing.T) {
	assert.Equal(t, "stringIsFun", ConvertCamelCase("string-is-fun"))
}

func TestConvertCamelCaseUnderscore(t *testing.T) {
	assert.Equal(t, "stringIsFun", ConvertCamelCase("string_is_fun"))
}
