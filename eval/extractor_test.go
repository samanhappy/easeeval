package eval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractorJson(t *testing.T) {
	v, e := ExtractJson(`{"k":2}`, ".k")
	assert.Nil(t, e)
	assert.Equal(t, "2", v)

	v, e = ExtractJson(`{
		"code": 0,
		"msg": "success",
		"data": [
			{
				"orderPayAt": "2022-07-28T20:55:06"
			}
		]
	}`, ".data.[0].orderPayAt")
	assert.Nil(t, e)
	assert.Equal(t, "2022-07-28T20:55:06", v)
}
