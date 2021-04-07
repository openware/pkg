package barong

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	"github.com/stretchr/testify/assert"
)

const (
	URL           = "http://barong:8080/api/v2/management"
	jwtIssuer     = "applogic"
	jwtAlgo       = "RS256"
	jwtPrivateKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBelNSeHpxZkhpVFg1bzl5N0JBdE1NM0lxcmtrSWNMZmhia3FHS0V3VFJXYkMyam5ZCmRDaXQ5SC9BV25zYTdpcnlOQ0hwS1lhUVozMkJ4MVVycnQrVk9kMm1YSDEwZHJ5VUtQcTZDdk1rSnBqYitNTncKTXlmd0dxKzdNdmMrUDBWcXp3dE5oNnplVThubVRzMTY2eWd0SVdzREEydWprc1R3ZHY3bEVFK0xMY1djbC90Ugp2dCtKcWtqeVNDYm1UOWl2bXUyeWh4UGFWbmU1TGxLQ2JnOWVJZEZTWEV1R2JFSnBpVGNhZ0lsSUh3VmJ6VnpSCllCbzlobjRXbHhXSmVUSkNQYmxIN0U0cmtrNUdJUDJqUnlvYmdUb2pTNERHS1hqc1ZwamJtc0l1Vk15OE5qQjUKWTJVcU54MVRaWFcvZTV2Nk9jUk9mOFFXSmxhQW1jNTd3S09La3Y0QXdoZ3QzWnluTXpnaVZaQ0Q2MzFDMmszYgpIUitTeHNneXFocHFZSEhUMThxM1hlVnI3d2lsR1Q1WnBremlIbk9SbjlFRjIxa21jczc2MWNBS3hIU3lENkNwCmdTVERObkc0Sm5sWjVnM1BpSk9YR1IvajgvZzNmZ3YydnAzL05nVm55Q2IrTUhCT0ZSaStXOXM1ZytMRU9XdTEKNmp1b2lWOHpaNlk5UmZFdHhVZjRzZ0ErNFJjUGZxQkJ5U2JPRW9neDJ5dDB3aVVJeUwyTEpZMmpBQUNxTVpSWApVUnB0bEtsVjAzbWZuL0I2aHkzNWR5VmN4ME9WRlpXK01tZUgzaHNHSWQwZE92UWdQaHpHOVBkcTVCUWVoQkp5ClpvdlU4RE54U0p5dnA1N0tqYlJWY2VIeklzUUw0MGZXdW81VlU0bHdnLzZueDdmYWlsdkRKSS9hZXQwQ0F3RUEKQVFLQ0FnQU1YRzdUSWY3L0FKYWJUaGlpeEwrQnRoWm1UQlpMSEhsaitPK2VpLzc1UnBqbEoya29qcTcwdGFIMAprY2hzbzMvV3JsaHJYU1ZrWndhajZUanBuNlZSU0U3VzhlUkxwMDlTTE5GN0NXMmJPY2kvYzU5V0pjanRBcnZICjlXZjF6Z3dDajg3TEp4cDZlQWI5cHBvS2czQTh2RU1CT01JeGZOWjBoU1Z1Vnl5dXhHS01NZU9hR2NRazA2SnQKd0pKT0syTmhkWU0xYW5mVWtBQkRqMHMyc0l4ZWcwdHdMa2phU3lJcTEzd3NWSmxZN1N5NzhpVFhvcDBrZG9LTAo5Z3REbDBpd2lYS1JCYURRZnhEd3VmZlZ1TzdSV1p4NDF6aVpsU1RBanhOa2Z1RGwwVFJpRzRlaytwcVJtWjNGCjFsT0Vja0NnckhpQ2NHRlpUQXNSdVlSeGRpbEtXSSthRU15L2lxUzFja0dPR0NuNUVMSmNpT21HV0hTQTF2eFgKZ1FMaVZ1Nkh1b0ZpcW1Jd3NWeHF5MFczN1ZDUEpVWFRXVmhCOGRXbjJscklyd3U0WWhhZzEwcXJ3cnJIQzFIWAo3a0hyU3JJRGdCSzBNdlQyTFh1TUU2eFcxT2k2N2k2RG05UHNrVXRIanVUNEpWTk94YUdRVUVBQm5YSkd5Y1MzCkRFVUFoR25qRmdpY29vM2JoQXZrdmh4WXNFdUdVRDZiVlZ1UTN0MytnbW0rWWVoSlUxdms4bldLYVdhOG14WEYKQWYrZmJIZ2c4TWxaVnJjOWpQQmZJZDQxdEZlUFpQbXpaaWFHRThnbHJ1d2xzQkhHa2F5WG4xYWtqM3JiV2NscgpLSWlRemJ4UjczTmZVTHNhZ2ZtL3lJWjBQeDRjWHgydG9zYUZiU3lkbGZLVDhaNGNEUUtDQVFFQTVyUVRwT3QzCjQ3VkZZUkw1ZTNaZTc3Rjc1MG9WN0pob3Q5YkZjenJFQTl6cVRFTC9Ua2pGbVE3RGR4aFgwTVo5YVZnMklMY3EKaThOdnpZWWhiMDlwd1ZsUWpzQmtocXA5WjlYWWU5bkNJRHBCMFhQWnRUM2pqbUtFcUYwZC8yT0FHeW16bWhVcwp6ajhnNDRsbFV5K1dvc1p2L2NicmF1RHZaUDBqbWd6cTRWSDlRK21VR2RKZlRDNXpBYldnZ3RodjhyWXd4YUVyCnhxOE9DNVp0V1BWK3BwN1dKV3ZSRWs4WWtTcFZVWjJVV2R0V2VaUk5hNWIxUHBUS3J4VGRxNkUvb3V4d1UwOTMKYWIwWDVlMGs3bnErVyttK1Y5Uks0NEUyL0x1c3JUWUlkRVRuWUpsRVo0aU9EUEJGNit5QlBtMVBKbGtSN3RjQwpNelhsUlNXRllUQjA0d0tDQVFFQTQ2TGNNUlZpdHd5QWkzamY1STd0amc4TmFpREw4dkJ1QzFzRXlMMndKaFNUClVXU0tiWmZIcEVpNlI1UVR2L01sd0RHNUZ4VThUWk1pSHl1RFFMR0pwYnQ3Vk5mZFVpWlozZkJiRkNKK3htaDMKOU1FRWszcWZJZlVHN0d3RmlkUEx4R3Q0OGMrREFNZW5nYnVhNlpMRTdvd21SZW8wK3k3cmRpQ3d3MEY3MWxUSApzbVd1aEhCa1hoTFlpU1lpZ1ZRTVJDR0U3aW5LdUdkQTdLd2dpTWQ4VVhrV3NibDJZRmFmd3JSRlVZVmJRUlJ2ClZTVnVMYVVoYTZCbzFkbGxLZ1lrVFIwOEZRNWhrVisvaGtZNlZoM1BCRzZEZXhJR1FiNlJrQjdDNWprSVhDdlMKUzFvOWF6aENreTFralo4YXpuMUg2Ky9xM25qckl5UXNtcWdoZUpBZFB3S0NBUUJBNG1Tai9aVzZkVUVPREVnZQpjU3hDUGFpYlpEckdVQmNqblVQckpKdjhlaVZyVFd5QWwvYjdGU3ZrVXZSZnczT0NMVTBMNW5nUTF1YWE1eDZBCkw5V09pNUFjbGYrdjRFTms4TC95RlV5RHc5Ni9DZFl4SXpiYzFOaDZnYlh1SGczcGxkRHRoUWNVK3F4RlVsOHQKQmpWWGtuZnM2QVZPQ2ZWS2NlZVJiQkNqVG12c3JjVDVmakZQTzhFY3VmaHExSFNuenBYby8ydFFkZXQ5VnRGcQpNNkZyTzBEL1JWT0gwcmNXSE5IaUltK1cxaGw4R0RtdUNNYncwdWd1VmJBQ2xWZFFleThjUHoxV2Y5ZzQwbm1RCm1QVHc1TXlqNXhFbzZ5Nkw1anlxZW9mbUszcm5zRE9NNnRzSXlJcmh6NktKN0RSV2xMWjJkZ0lvWlFBV2NuY1EKM3BBQkFvSUJBQmQ2Q1dtS2doYk0xRWtPRzFFd0tISFpQWkh2ZGZsRk1LUTlLOTRrS2hHVFY2b3lTMUNJTWMvUQpyRjJMZVFuMzRySFNydnNoZG9tdG5meEcrWTluZ0FHMnR6NkYwTTZUSS91T3VXWDNOTW56cGtONDBLY0JJMzVXCkRmTytKRWdWcnROQUhrWWFGN0d4NWFXc21vcHlWNXNlbXlma3dyZ1JHN21nSDNyVHV4amN2NGUza3VzWHlGSW4KY1d1Ym9qMWlWSzJHSTNhSW10NnZ6M05aUVRXNkZTazE2dEJEaDJEaUxqSGZjN0szcFRTdURkbGpOZHpCUmhRYQpoQlZpQ1Z2dkxEbER4Wm1LVlNld0QwbWkzb3RaSWF1Y1ZqVVFJOU1OKzJjNHRQTVhlTFJBMUx4dXZ4emF2WXIrClNIdU9xQzRabjV4R3J4dG9yeDk5c0pmMnRSVUJEL01DZ2dFQkFMRExRNzg0SDU2QjhlZU5zZE1wUk1mTitxZDQKclFlVTB2NUxUOElxdldTcmpLZTlHS3J4NTNSZ3RPNEtRTkdYc3FoMkROdktobFNDMys0amxCdms2UmxxUFdFYwpOSnRZSEw4UFBwby9hOEphM0F1RlpYL0NLNFFCeHR2SEdodjdPeG9JMG9zd0EzNG5NcVk2QXJaSktuSjEvcDljCmYxQk03TGRZQk1VNmVEeDBsSDVFM2xkM2lXVFN1ZUdWVk5PdzBpNmpoeDl3MUp0LzZwRis5NDJqdDFiRUoyN3YKYVdXT2REQ1g0SVIxMStiRlhhOEZJcEhCbStoTm1FdWRRc2hwN2pId2hCTjNiZnNSeHJXWGUyd1cvYkthdFBqWAo1N0p1bEFQVlN3L3h1TGJZZFZiVGlvdmRsMWxObXFJZEpqYVZma2ZZSzVJUVR1R0pxVHNzdVkvbWNITT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"
)

// Mock client
type MockClient struct {
	response []byte
	apiError *mngapi.APIError
}

// Mock request function
func (m *MockClient) Request(method string, path string, body interface{}) ([]byte, *mngapi.APIError) {
	return m.response, m.apiError
}

func TestCreateNewClient(t *testing.T) {
	t.Run("Success creation", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)

		assert.NotNil(t, client)
		assert.Nil(t, err)
	})

	t.Run("JWT issuer unset", func(t *testing.T) {
		client, err := New(URL, "", jwtAlgo, jwtPrivateKey)

		assert.Nil(t, client)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "JWT issuer unset")
	})

	t.Run("Invalid signing algorithm", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, "RS999", jwtPrivateKey)

		assert.Nil(t, client)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Unsupported signing method RS999")
	})

	t.Run("Invalid private key", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, "")

		assert.Nil(t, client)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "Invalid Key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	})
}

func TestCreateServiceAccount(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":"test+SI0388B7681C@yellow.com","uid":"SI0388B7681C","role":"service_account","level":3,"state":"active","user":{"email":"test@test.com","uid":"IDCA2AC08296","role":"superadmin","level":3,"otp":true,"state":"active","referral_uid":"","data":"{\"onboarding\":true,\"language\":\"en\"}"},"created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "IDCA2AC08296",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(serviceAccount)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Success response w/ state", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":"test+SI0388B7681C@yellow.com","uid":"SI0388B7681C","role":"service_account","level":3,"state":"active","user":{"email":"test@test.com","uid":"IDCA2AC08296","role":"superadmin","level":3,"otp":true,"state":"active","referral_uid":"","data":"{\"onboarding\":true,\"language\":\"en\"}"},"created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "IDCA2AC08296",
			Role:     "service_account",
			State:    "active",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(serviceAccount)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error user doesn't exist", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "User doesnt exist",
			},
		}

		params := CreateServiceAccountParams{
			OwnerUID: "ID123456789",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "User doesnt exist")
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":test@test.com}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "ID123456789",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateServiceAccountParams{
			OwnerUID: "ID123456789",
			Role:     "service_account",
			Level:    2,
		}
		serviceAccount, apiError := client.CreateServiceAccount(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})
}

func TestDeleteServiceAccountByUID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":"test+SI0388B7681C@yellow.com","uid":"SI0388B7681C","role":"service_account","level":3,"state":"active","user":{"email":"test@test.com","uid":"IDCA2AC08296","role":"superadmin","level":3,"otp":true,"state":"disabled","referral_uid":"","data":"{\"onboarding\":true,\"language\":\"en\"}"},"created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		uid := "SI0388B7681C"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)
		assert.Nil(t, apiError)

		result, err := json.Marshal(serviceAccount)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record is not found", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Record is not found",
			},
		}

		uid := "SI123456789"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Record is not found")
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"email":123}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		uid := "SI123456789"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{error}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		uid := "SI123456789"
		serviceAccount, apiError := client.DeleteServiceAccountByUID(uid)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, serviceAccount)
	})
}

func TestCreateAPIKeys(t *testing.T) {
	t.Run("Successful response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"kid":"133742","algorithm":"HS256","scope":["trade"],"state":"active","secret":"something_in_the_way","created_at":"2021-02-15T10:15:18Z","updated_at":"2021-02-15T10:15:18Z"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateAPIKeyParams{
			UID:       "IDCA2AC08296",
			Algorithm: "HS256",
			Scopes:    "trade",
		}
		apiKey, apiError := client.CreateAPIKey(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(apiKey)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error could not save secret", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "api_key.could_not_save_secret",
			},
		}

		params := CreateAPIKeyParams{
			UID:       "IDCA2AC08296",
			Algorithm: "HS256",
			Scopes:    "trade",
		}
		apiKey, apiError := client.CreateAPIKey(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "api_key.could_not_save_secret")
		assert.Nil(t, apiKey)
	})
}

func TestCreateAttachment(t *testing.T) {
	t.Run("Successful response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"user_uid":"UID388B768"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateAttachmentParams{
			UID:      "UID388B768",
			FileName: "test",
			FileExt:  ".png",
			Upload:   "iVBORw0KGgoAAAANSUhEUgAAAAUAAAAFCAYAAACNbyblAAAAHElEQVQI12P4//8/w38GIAXDIBKE0DHxgljNBAAO9TXL0Y4OHwAAAABJRU5ErkJggg==",
		}
		attachment, apiError := client.CreateAttachment(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(attachment)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error could not save secret", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "attachment.could_not_save",
			},
		}

		params := CreateAttachmentParams{
			UID:      "IDCA2AC08296",
			FileName: "test",
			FileExt:  ".png",
			Upload:   "iVBORw0K",
		}
		apiKey, apiError := client.CreateAttachment(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "attachment.could_not_save")
		assert.Nil(t, apiKey)
	})
}
