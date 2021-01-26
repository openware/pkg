package mgntapi

import (
	"testing"
)

func TestAuth_ManagementAPIV2(t *testing.T) {
	// assert := assert.New(t)

	t.Run("", func(t *testing.T) {
		client := New(
			"https://dev.yellow.openware.work",
			"/api/v2/barong/management/",
		)

		res, err := client.Request("GET", "xxx", []byte{})
		t.Error(res, err)
	})
}
