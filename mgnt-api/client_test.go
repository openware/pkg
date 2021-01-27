package mgntapi_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/openware/pkg/ika"
	mgntapi "github.com/openware/pkg/mgnt-api"
	"github.com/stretchr/testify/assert"
)

func Test_ManagementAPIV2_ValidConfiguration(t *testing.T) {
	assert := assert.New(t)

	const encodedKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBNW5hUThndjh1MlIvUE9STnRuVVhWSXlpMmdaWUdCQ0tmZDgxT29ZK2tSNDYxdmZ1ClIwZWo2RC9LTFlRbEViRWNWbzIrVnkxa01DaWNXdDZFSmhzUEk4YTN0K1djMWhQc05wVkJhRWhoWWdQS2hOQ2YKb2JnZ0puTXRVU2dRdm1BQnQrR2NPQlBOeWxJdlZrT21hTVZTZHd4anFzNTBtMVdJenVHb2dFcFBqY2pMTW0xMApQcE5WNHF6UVN0ZVdXbnUzNUg2R3N6M3BtSGJkdm90UjZYeEM2WDdIYmxvRXc1R0Y1dHB2VmR4MlhBVC9IbHNnCi8xN0hyRTlCRmhGODRONzh4Qlc5SGpVbnZNeGE5d3pUaERSQ0ZlakJUZ1VJSWpKSzBwUXNibFFvL3g2dDMrUzUKV3BaNDU4ZzJvU1hnaXkveW1JTFhsUFIxeU1BN2lNa1hWVGpPbVFJREFRQUJBb0lCQVFDOSt4V25ncnczbWpQTgpkWUtlbTAyOU5DWDNSdTJPQU95NXNLd0hiNnphSWlwdEZYc0dwWWIzcU1ZNDJVdFpsei8rRmVESHFyS0JoS2pICnU5RUNQS0l4WXRvR0xiRXBSTWtmZ2RDbWI2eGZpVEtFWkJxRHpPNHI1QnlDWDEzV0lmeW9vY0lPOUR4YndYNG0KUmFSRGtBNVg5dzJlTzQwaWs5TXdnQk5RbG5HWU5WSnhLYjl1V1VWNm9RdmVSNUppeVNqVDlXVGp3VThUTUEvRwo1NjVzUmxVN3BqSGF6d3diQk5WaDV6OWx3clpUNHplaFhqcmZXcERNMzFHTDkzcFlMdUh6S3VzUVpCbytpTnlUCitRNkRQSTQ4Tzl5QkdFdG5jaHZwWFBnZi9pWWdTeGgyM3pwWmgxVUVWMDN6WkVHY2lGY3FTYUJFb2ZTM3p2bUMKMVJUNHd4RlZBb0dCQU9nQXJOajJyZGxuckJrUDJxcWdVSVBZYkpZVjBraFU1c20zMHhybjNkcDNBejNhQzdVVgpKVnJpUFl2ZktqSWNXUFVLMkVqYTZHRFRGWUttTWhjK1lWaFlwZjZwdmVXZUhmNS8xYkVra3l2NlZ1d21DTGRSCmU2UTVtTU5aK09wV2dHQTZuWkM1TDVSeitJbFFlM1cxR0NIaDMwRkZpaFBQdlRxN0JJSHgyamZYQW9HQkFQNU4KSUVVd3dtK2dFZWpaMTBlRlNlcG8yWXJkQjUwK1IxZDVrZ0J0amZleXczL1ZCVXpWU2ROdFhxSFBvUDZDT09zcwpZbWhtWmJZOGI1cmZjM3JZbUMxd2MxUjNaZklDL2thNHBaN2tSbXF4RnQwbFJYakVaN0RQeUsvT3ArVWFUeXpxCnFBK2FSWnhwNzBEWXNiVytoMGxLMGJESnVkS2gvV0czcTFQNUo1OFBBb0dCQU1qNGhNSnhkWW1saitRcDRxNzUKcnFWM25pQ3BDSDZWNVZJS0Jqb0JieUluQkV6WkRGa3gxeWtTWUdSQXppbVllc3JTT1NkclVlOUdDeFVnNkxWUAoyVDJSbFVHMFFvYWM0TGlzZmkwMFZMUzg1LzBxdVZRcnBxSk5MbkxUQnBmZ2xOWkhFR1RrdGoydjlEVG0zZnZLCkF2eWUwQ21YbHBPdzJlZjlSMXRWYVVZREFvR0JBTE5WbnFGTXJvSGJ4MldIWW1zY2t6RE0ra1VVZEk4dVlVOU4KKzJsejJQOUtRTlpBV25tQm5JdU9nSUxxRW1ZSlhheHpZMzZ1WDZJeFlwODhYNHJOZmh2bFJsL1Z4NzN3NEhMdApPbGNnTW95QkVGZXFOaURobVNJMmxoZHRURGVqNHh3UTY2MzlKSVFXck5QMVBQV25SRzZxWmRBZm9uenBJZkFzCmY1VTdpdlovQW9HQUhZRE9HZ3hoeklHRWpFcFMyNnZSSUJsVVlDb0IrcHY1Um5rbXN6RFE2Z3cwaE8zb1U3bGEKN1BGNHFHQzgrSlRwZ1JNTVZVYlJuSzZYVDF0cWY3dFJKZnF6aXRnZnhsa0N0Nm9zekUvT3BPaWZ5VVY0YjhwOQpuZUMyaE1TVDlENDBRVy9mcnZxZTFwVDQvaW5vUWREU3grS29JREtDTkEvWHdUMWJDZXZ6NUU4PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo="
	configFile := fmt.Sprintf(`
barong:
  actions:
    otp_sign:
      required_signatures:
      - applogic
      requires_barong_totp: false
    write_documents:
      required_signatures:
      - applogic
      requires_barong_totp: false
    write_labels:
      required_signatures:
      - applogic
      requires_barong_totp: false
    write_phones:
      required_signatures:
      - applogic
      requires_barong_totp: false
  jwt: {}
  keychain:
    applogic:
      algorithm: RS256
      value: "%s"
peatio:
  actions:
    read_withdraws:
      required_signatures:
      - applogic
      requires_barong_totp: false
    write_withdraws:
      required_signatures:
      - applogic
      requires_barong_totp: false
  jwt: {}
  keychain:
    applogic:
      algorithm: RS256
      value: "%s"
`, encodedKey, encodedKey)

	tmpFile, err := ioutil.TempFile(os.TempDir(), "*.yml")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpFile.Name())

	input := []byte(configFile)
	if _, err = tmpFile.Write(input); err != nil {
		t.Error(err)
	}

	cfg := mgntapi.ManagementAPIV2Config{}
	ika.ReadConfig(tmpFile.Name(), &cfg)

	t.Run("valid barong configuration", func(t *testing.T) {
		// validate algorithm and value
		assert.Equal(cfg.Barong.Keychain.Applogic.Algorithm, "RS256")
		assert.Equal(cfg.Barong.Keychain.Applogic.Value, encodedKey)

		// validate actions
		assert.Equal(len(cfg.Barong.Actions.OTPSign.Signatures), 1)
		assert.Equal(cfg.Barong.Actions.OTPSign.Signatures[0], "applogic")
		assert.Equal(cfg.Barong.Actions.OTPSign.RequireTOTP, false)
		assert.Equal(len(cfg.Barong.Actions.WriteDocuments.Signatures), 1)
		assert.Equal(cfg.Barong.Actions.WriteDocuments.Signatures[0], "applogic")
		assert.Equal(cfg.Barong.Actions.WriteDocuments.RequireTOTP, false)
		assert.Equal(len(cfg.Barong.Actions.WriteLabels.Signatures), 1)
		assert.Equal(cfg.Barong.Actions.WriteLabels.Signatures[0], "applogic")
		assert.Equal(cfg.Barong.Actions.WriteLabels.RequireTOTP, false)
		assert.Equal(len(cfg.Barong.Actions.WritePhones.Signatures), 1)
		assert.Equal(cfg.Barong.Actions.WritePhones.Signatures[0], "applogic")
		assert.Equal(cfg.Barong.Actions.WritePhones.RequireTOTP, false)

		// validate missing actions
		assert.Equal(len(cfg.Barong.Actions.ReadUsers.Signatures), 0)
		assert.Equal(len(cfg.Barong.Actions.WriteUsers.Signatures), 0)
		assert.Equal(len(cfg.Barong.Actions.ReadDocuments.Signatures), 0)
		assert.Equal(len(cfg.Barong.Actions.ReadLabels.Signatures), 0)
		assert.Equal(len(cfg.Barong.Actions.ReadPhones.Signatures), 0)
	})

	t.Run("valid peatio configuration", func(t *testing.T) {
		// validate algorithm and value
		assert.Equal(cfg.Peatio.Keychain.Applogic.Algorithm, "RS256")
		assert.Equal(cfg.Peatio.Keychain.Applogic.Value, encodedKey)

		// validate actions
		assert.Equal(len(cfg.Peatio.Actions.ReadWithdraws.Signatures), 1)
		assert.Equal(cfg.Peatio.Actions.ReadWithdraws.Signatures[0], "applogic")
		assert.Equal(cfg.Peatio.Actions.ReadWithdraws.RequireTOTP, false)
		assert.Equal(len(cfg.Peatio.Actions.WriteWithdraws.Signatures), 1)
		assert.Equal(cfg.Peatio.Actions.WriteWithdraws.Signatures[0], "applogic")
		assert.Equal(cfg.Peatio.Actions.WriteWithdraws.RequireTOTP, false)

		// validate missing actions
		assert.Equal(len(cfg.Peatio.Actions.ReadAccounts.Signatures), 0)
		assert.Equal(len(cfg.Peatio.Actions.WriteAccounts.Signatures), 0)
		assert.Equal(len(cfg.Peatio.Actions.ReadOrders.Signatures), 0)
		assert.Equal(len(cfg.Peatio.Actions.WriteOrders.Signatures), 0)
		assert.Equal(len(cfg.Peatio.Actions.ReadMarkets.Signatures), 0)
		assert.Equal(len(cfg.Peatio.Actions.WriteMarkets.Signatures), 0)
	})
}
