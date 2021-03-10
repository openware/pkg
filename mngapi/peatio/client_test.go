package peatio

import (
	"encoding/json"
	"testing"

	"github.com/openware/pkg/mngapi"
	"github.com/stretchr/testify/assert"
)

const (
	URL           = "http://peatio:8080/api/v2/management"
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
		assert.EqualError(t, err, "Invalid Key: Key must be PEM encoded PKCS1 or PKCS8 private key")
	})
}

func TestGetCurrencyByCode(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"bnb","name":"Binance Coin","descritpion":"","homepage":"","price":"23.8","explorer_transaction":"https://kovan.etherscan.io/tx/#{txid}","explorer_address":"https://kovan.etherscan.io/address/#{address}","type":"coin","deposit_enabled":true,"withdrawal_enabled":true,"deposit_fee":"0.0","min_deposit_amount":"0.3455425","withdraw_fee":"0.0","min_withdraw_amount":"0.3455425","withdraw_limit_24h":"100000.0","withdraw_limit_72h":"200000.0","base_factor":1000000000000000000,"precision":10,"position":48,"icon_url":"https://sorage.googleapis.com/devel-yellow-exchange-applogic/uploads/asset/icon/bnb/8ea0f30c1b.png","min_confirmations":10,"code":"bnb","min_collection_amount":"0.3455425","visible":true,"subunits":18,"options":{"erc20_contract_address":"0xb8c77482e45f1f44de1745f52c74426c631bdd52"},"created_at":"2020-02-24T15:34:03+01:00","updated_at":"2020-12-02T10:42:33+01:00"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		currency, apiError := client.GetCurrencyByCode("bnb")
		assert.Nil(t, apiError)

		result, err := json.Marshal(currency)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "Invalid",
			},
		}

		currency, apiError := client.GetCurrencyByCode("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Invalid")
		assert.Nil(t, currency)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":bnb,price:123.456}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		currency, apiError := client.GetCurrencyByCode("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{""}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		currency, apiError := client.GetCurrencyByCode("bnb")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, currency)
	})
}

func TestCreateWithdraw(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":"TIDE54B7D229E","uid":"ID16421C020A","currency":"btc","note":"","type":"coin","amount":"0.1195","fee":"0.0005","rid":"1CzSHQnuwp52ErrrtM169FW4FuuRhEksMR","state":"skipped","created_at":"2021-01-12T07:27:41+01:00","blockchain_txid":"","transfer_type":"crypto"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(withdraw)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, withdraw)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateWithdrawParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
			Amount:   10.0,
		}
		withdraw, apiError := client.CreateWithdraw(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})
}

func TestGetWithdrawByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":"TIDE54B7D229E","uid":"ID16421C020A","currency":"btc","note":"","type":"coin","amount":"0.1195","fee":"0.0005","rid":"1CzSHQnuwp52ErrrtM169FW4FuuRhEksMR","state":"skipped","created_at":"2021-01-12T07:27:41+01:00","blockchain_txid":"","transfer_type":"crypto"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		withdraw, apiError := client.GetWithdrawByID("TIDE54B7D229E")
		assert.Nil(t, apiError)

		result, err := json.Marshal(withdraw)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record not found", func(t *testing.T) {
		client, _ := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Couldn't find record.",
			},
		}

		withdraw, apiError := client.GetWithdrawByID("TIDXXXX")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, withdraw)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		withdraw, apiError := client.GetWithdrawByID("TIDE54B7D229E")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{","}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		withdraw, apiError := client.GetWithdrawByID("TIDE54B7D229E")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, withdraw)
	})
}

func TestGetAccountBalance(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","balance":"996.23352165725","locked":"0.0"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(balance)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record not found", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Couldn't find record.",
			},
		}

		params := GetAccountBalanceParams{
			UID:      "ID1234567890",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, balance)
	})

	t.Run("Error invalid currency", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "currency does not have a valid value",
			},
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnbxxx",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "currency does not have a valid value")
		assert.Nil(t, balance)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"balance":996.23352165725,locked:0.0}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, balance)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{aaa: 1}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetAccountBalanceParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		balance, apiError := client.GetAccountBalance(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, balance)
	})
}

func TestGenerateDepositAddress(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","address":"0x5b89a2a38b7398c71cfc420a6ed3b5f2a1a01a3e","currencies":["usdt","bnb","uni"],"state":"active","remote":false}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(paymentAddress)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Couldn't find record.",
			},
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, paymentAddress)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","address":0x5b89a2a38b7398c71cfc420a6ed3b5f2a1a01a3e,"currencies":["usdt","bnb","uni"],"state":"active","remote":false}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, paymentAddress)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{[]}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GenerateDepositAddressParams{
			UID:      "IDCA2AC08296",
			Currency: "bnb",
		}
		paymentAddress, apiError := client.GenerateDepositAddress(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, paymentAddress)
	})
}

func TestCreateDeposit(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"tid":"TIDBD6B265303","currency":"usd","address":"","uid":"ID732785AC58","type":"fiat","amount":"750.77","state":"submitted","created_at":"2021-03-02T07:33:02+01:00","completed_at":null,"transfer_type":"fiat"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "usd",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(deposit)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "bnb",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, deposit)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "bnb",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateDepositParams{
			UID:      "ID732785AC58",
			Currency: "bnb",
			Amount:   10.0,
		}
		deposit, apiError := client.CreateDeposit(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})
}

func TestGetDepositByID(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"tid":"TIDF6289303E1","currency":"btc","address":"","uid":"ID6CBD4E84C7","type":"coin","amount":"6346.0","state":"submitted","created_at":"2021-03-02T05:54:52+01:00","completed_at":null,"blockchain_txid":"56bzwdd359kxd0r3qt3mz1cbcrc8o3r5hshlgbag42z7ka2o9hd4b5me80hh0khb","blockchain_confirmations":711753,"transfer_type":"crypto"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		deposit, apiError := client.GetDepositByID("TIDF6289303E1")
		assert.Nil(t, apiError)

		result, err := json.Marshal(deposit)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error record not found", func(t *testing.T) {
		client, _ := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "Couldn't find record.",
			},
		}

		deposit, apiError := client.GetDepositByID("TIDXXXX")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "Couldn't find record.")
		assert.Nil(t, deposit)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"tid":1234}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		deposit, apiError := client.GetDepositByID("TIDF6289303E1")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{","}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		deposit, apiError := client.GetDepositByID("TIDF6289303E1")

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposit)
	})
}

func TestGetDeposits(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"tid":"TID9119EEAE36","currency":"usd","address":"","uid":"ID9C5C7208EB","type":"fiat","amount":"8423.0","state":"collected","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","transfer_type":"fiat"},{"id":2,"tid":"TID17505F194C","currency":"btc","address":"","uid":"ID0B0C77487A","type":"coin","amount":"191.0","state":"fee_processing","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","blockchain_txid":"wfmvae8elj0egr309u9oodl58ypzifdfjz9vd1i82t3ng4uepmokagack0shfsif","blockchain_confirmations":367597,"transfer_type":"crypto"}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)
		assert.Nil(t, apiError)
		assert.NotNil(t, deposits)

		result, err := json.Marshal(deposits)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 422,
				Error:      "Error",
			},
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 422)
		assert.Equal(t, apiError.Error, "Error")
		assert.Nil(t, deposits)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"tid":"TID9119EEAE36","currency":"usd","address":"","uid":"ID9C5C7208EB","type":"fiat","amount":8423.0,"state":"collected","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","transfer_type":"fiat"},{"id":2,"tid":"TID17505F194C","currency":"btc","address":"","uid":"ID0B0C77487A","type":"coin","amount":"191.0","state":"fee_processing","created_at":"2021-03-02T04:40:06+01:00","completed_at":"2021-03-02T04:40:06+01:00","blockchain_txid":"wfmvae8elj0egr309u9oodl58ypzifdfjz9vd1i82t3ng4uepmokagack0shfsif","blockchain_confirmations":"367597","transfer_type":"crypto"}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposits)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetDepositsParams{
			UID: "IDCA2AC08296",
		}
		deposits, apiError := client.GetDeposits(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, deposits)
	})
}

func TestCreateEngine(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"opendax_cloud","driver":"opendax","uid":"UID123123","url":"https://example.com","state":"online"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.CreateEngine(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(engine)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.CreateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, engine)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": opendax_cloud}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.CreateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := CreateEngineParams{
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.CreateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})
}

func TestUpdateEngine(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":1,"name":"opendax_cloud","driver":"opendax","uid":"UID123123","url":"https://example.com","state":"online"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.UpdateEngine(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(engine)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}

		engine, apiError := client.UpdateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, engine)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": opendax_cloud}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.UpdateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateEngineParams{
			ID:     "1",
			Name:   "opendax_cloud",
			Driver: "opendax",
			UID:    "UID123123",
			URL:    "https://example.com",
			State:  1,
			Key:    "key",
			Secret: "secret",
		}
		engine, apiError := client.UpdateEngine(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})
}

func TestGetEngines(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":1,"name":"opendax_cloud","driver":"opendax","uid":"UID123123","url":"https://example.com","state":"online"}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(engine)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, engine)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"name": opendax_cloud}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"-"}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := GetEngineParams{
			Name: "opendax_cloud",
		}
		engine, apiError := client.GetEngines(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, engine)
	})
}

func TestGetMarkets(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"id":"btcusd","name":"BTC/USD","base_unit":"btc","quote_unit":"usd","min_price":"0.01","max_price":"0.0","min_amount":"0.00000001","amount_precision":8,"price_precision":2,"state":"enabled","position":1,"engine_id":1,"created_at":"2021-03-05T14:52:48+01:00","updated_at":"2021-03-05T14:52:48+01:00"}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		markets, apiError := client.GetMarkets()
		assert.Nil(t, apiError)

		result, err := json.Marshal(markets)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		markets, apiError := client.GetMarkets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, markets)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"name": BTC/USD}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		markets, apiError := client.GetMarkets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, markets)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `[{"-"}]`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		markets, apiError := client.GetMarkets()

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, markets)
	})
}

func TestUpdateMarket(t *testing.T) {
	t.Run("Success response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"id":"btcusd","name":"BTC/USD","base_unit":"btc","quote_unit":"usd","min_price":"0.01","max_price":"0.0","min_amount":"0.00000001","amount_precision":8,"price_precision":2,"state":"enabled","position":1,"engine_id":1,"created_at":"2021-03-05T14:52:48+01:00","updated_at":"2021-03-05T14:52:48+01:00"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}

		market, apiError := client.UpdateMarket(params)
		assert.Nil(t, apiError)

		result, err := json.Marshal(market)
		assert.NoError(t, err)
		assert.Equal(t, result, []byte(expected))
	})

	t.Run("Error response", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		client.mngapiClient = &MockClient{
			response: nil,
			apiError: &mngapi.APIError{
				StatusCode: 404,
				Error:      "404 Not Found",
			},
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}

		market, apiError := client.UpdateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 404)
		assert.Equal(t, apiError.Error, "404 Not Found")
		assert.Nil(t, market)
	})

	t.Run("Error mismatch data type during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"name": BTC/USD}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}
		market, apiError := client.UpdateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})

	t.Run("Error invalid json response during unmarshal", func(t *testing.T) {
		client, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		assert.NoError(t, err)

		expected := `{"-"}`
		client.mngapiClient = &MockClient{
			response: []byte(expected),
			apiError: nil,
		}

		params := UpdateMarketParams{
			ID:       "1",
			EngineID: "1",
		}
		market, apiError := client.UpdateMarket(params)

		assert.NotNil(t, apiError)
		assert.Equal(t, apiError.StatusCode, 500)
		assert.NotEmpty(t, apiError.Error)
		assert.Nil(t, market)
	})
}
