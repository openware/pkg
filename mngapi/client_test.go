package mngapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	jwtgo "github.com/golang-jwt/jwt"
)

const (
	URL           = "http://peatio:8080/api/v2/management/"
	jwtIssuer     = "applogic"
	jwtAlgo       = "RS256"
	jwtPrivateKey = "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBelNSeHpxZkhpVFg1bzl5N0JBdE1NM0lxcmtrSWNMZmhia3FHS0V3VFJXYkMyam5ZCmRDaXQ5SC9BV25zYTdpcnlOQ0hwS1lhUVozMkJ4MVVycnQrVk9kMm1YSDEwZHJ5VUtQcTZDdk1rSnBqYitNTncKTXlmd0dxKzdNdmMrUDBWcXp3dE5oNnplVThubVRzMTY2eWd0SVdzREEydWprc1R3ZHY3bEVFK0xMY1djbC90Ugp2dCtKcWtqeVNDYm1UOWl2bXUyeWh4UGFWbmU1TGxLQ2JnOWVJZEZTWEV1R2JFSnBpVGNhZ0lsSUh3VmJ6VnpSCllCbzlobjRXbHhXSmVUSkNQYmxIN0U0cmtrNUdJUDJqUnlvYmdUb2pTNERHS1hqc1ZwamJtc0l1Vk15OE5qQjUKWTJVcU54MVRaWFcvZTV2Nk9jUk9mOFFXSmxhQW1jNTd3S09La3Y0QXdoZ3QzWnluTXpnaVZaQ0Q2MzFDMmszYgpIUitTeHNneXFocHFZSEhUMThxM1hlVnI3d2lsR1Q1WnBremlIbk9SbjlFRjIxa21jczc2MWNBS3hIU3lENkNwCmdTVERObkc0Sm5sWjVnM1BpSk9YR1IvajgvZzNmZ3YydnAzL05nVm55Q2IrTUhCT0ZSaStXOXM1ZytMRU9XdTEKNmp1b2lWOHpaNlk5UmZFdHhVZjRzZ0ErNFJjUGZxQkJ5U2JPRW9neDJ5dDB3aVVJeUwyTEpZMmpBQUNxTVpSWApVUnB0bEtsVjAzbWZuL0I2aHkzNWR5VmN4ME9WRlpXK01tZUgzaHNHSWQwZE92UWdQaHpHOVBkcTVCUWVoQkp5ClpvdlU4RE54U0p5dnA1N0tqYlJWY2VIeklzUUw0MGZXdW81VlU0bHdnLzZueDdmYWlsdkRKSS9hZXQwQ0F3RUEKQVFLQ0FnQU1YRzdUSWY3L0FKYWJUaGlpeEwrQnRoWm1UQlpMSEhsaitPK2VpLzc1UnBqbEoya29qcTcwdGFIMAprY2hzbzMvV3JsaHJYU1ZrWndhajZUanBuNlZSU0U3VzhlUkxwMDlTTE5GN0NXMmJPY2kvYzU5V0pjanRBcnZICjlXZjF6Z3dDajg3TEp4cDZlQWI5cHBvS2czQTh2RU1CT01JeGZOWjBoU1Z1Vnl5dXhHS01NZU9hR2NRazA2SnQKd0pKT0syTmhkWU0xYW5mVWtBQkRqMHMyc0l4ZWcwdHdMa2phU3lJcTEzd3NWSmxZN1N5NzhpVFhvcDBrZG9LTAo5Z3REbDBpd2lYS1JCYURRZnhEd3VmZlZ1TzdSV1p4NDF6aVpsU1RBanhOa2Z1RGwwVFJpRzRlaytwcVJtWjNGCjFsT0Vja0NnckhpQ2NHRlpUQXNSdVlSeGRpbEtXSSthRU15L2lxUzFja0dPR0NuNUVMSmNpT21HV0hTQTF2eFgKZ1FMaVZ1Nkh1b0ZpcW1Jd3NWeHF5MFczN1ZDUEpVWFRXVmhCOGRXbjJscklyd3U0WWhhZzEwcXJ3cnJIQzFIWAo3a0hyU3JJRGdCSzBNdlQyTFh1TUU2eFcxT2k2N2k2RG05UHNrVXRIanVUNEpWTk94YUdRVUVBQm5YSkd5Y1MzCkRFVUFoR25qRmdpY29vM2JoQXZrdmh4WXNFdUdVRDZiVlZ1UTN0MytnbW0rWWVoSlUxdms4bldLYVdhOG14WEYKQWYrZmJIZ2c4TWxaVnJjOWpQQmZJZDQxdEZlUFpQbXpaaWFHRThnbHJ1d2xzQkhHa2F5WG4xYWtqM3JiV2NscgpLSWlRemJ4UjczTmZVTHNhZ2ZtL3lJWjBQeDRjWHgydG9zYUZiU3lkbGZLVDhaNGNEUUtDQVFFQTVyUVRwT3QzCjQ3VkZZUkw1ZTNaZTc3Rjc1MG9WN0pob3Q5YkZjenJFQTl6cVRFTC9Ua2pGbVE3RGR4aFgwTVo5YVZnMklMY3EKaThOdnpZWWhiMDlwd1ZsUWpzQmtocXA5WjlYWWU5bkNJRHBCMFhQWnRUM2pqbUtFcUYwZC8yT0FHeW16bWhVcwp6ajhnNDRsbFV5K1dvc1p2L2NicmF1RHZaUDBqbWd6cTRWSDlRK21VR2RKZlRDNXpBYldnZ3RodjhyWXd4YUVyCnhxOE9DNVp0V1BWK3BwN1dKV3ZSRWs4WWtTcFZVWjJVV2R0V2VaUk5hNWIxUHBUS3J4VGRxNkUvb3V4d1UwOTMKYWIwWDVlMGs3bnErVyttK1Y5Uks0NEUyL0x1c3JUWUlkRVRuWUpsRVo0aU9EUEJGNit5QlBtMVBKbGtSN3RjQwpNelhsUlNXRllUQjA0d0tDQVFFQTQ2TGNNUlZpdHd5QWkzamY1STd0amc4TmFpREw4dkJ1QzFzRXlMMndKaFNUClVXU0tiWmZIcEVpNlI1UVR2L01sd0RHNUZ4VThUWk1pSHl1RFFMR0pwYnQ3Vk5mZFVpWlozZkJiRkNKK3htaDMKOU1FRWszcWZJZlVHN0d3RmlkUEx4R3Q0OGMrREFNZW5nYnVhNlpMRTdvd21SZW8wK3k3cmRpQ3d3MEY3MWxUSApzbVd1aEhCa1hoTFlpU1lpZ1ZRTVJDR0U3aW5LdUdkQTdLd2dpTWQ4VVhrV3NibDJZRmFmd3JSRlVZVmJRUlJ2ClZTVnVMYVVoYTZCbzFkbGxLZ1lrVFIwOEZRNWhrVisvaGtZNlZoM1BCRzZEZXhJR1FiNlJrQjdDNWprSVhDdlMKUzFvOWF6aENreTFralo4YXpuMUg2Ky9xM25qckl5UXNtcWdoZUpBZFB3S0NBUUJBNG1Tai9aVzZkVUVPREVnZQpjU3hDUGFpYlpEckdVQmNqblVQckpKdjhlaVZyVFd5QWwvYjdGU3ZrVXZSZnczT0NMVTBMNW5nUTF1YWE1eDZBCkw5V09pNUFjbGYrdjRFTms4TC95RlV5RHc5Ni9DZFl4SXpiYzFOaDZnYlh1SGczcGxkRHRoUWNVK3F4RlVsOHQKQmpWWGtuZnM2QVZPQ2ZWS2NlZVJiQkNqVG12c3JjVDVmakZQTzhFY3VmaHExSFNuenBYby8ydFFkZXQ5VnRGcQpNNkZyTzBEL1JWT0gwcmNXSE5IaUltK1cxaGw4R0RtdUNNYncwdWd1VmJBQ2xWZFFleThjUHoxV2Y5ZzQwbm1RCm1QVHc1TXlqNXhFbzZ5Nkw1anlxZW9mbUszcm5zRE9NNnRzSXlJcmh6NktKN0RSV2xMWjJkZ0lvWlFBV2NuY1EKM3BBQkFvSUJBQmQ2Q1dtS2doYk0xRWtPRzFFd0tISFpQWkh2ZGZsRk1LUTlLOTRrS2hHVFY2b3lTMUNJTWMvUQpyRjJMZVFuMzRySFNydnNoZG9tdG5meEcrWTluZ0FHMnR6NkYwTTZUSS91T3VXWDNOTW56cGtONDBLY0JJMzVXCkRmTytKRWdWcnROQUhrWWFGN0d4NWFXc21vcHlWNXNlbXlma3dyZ1JHN21nSDNyVHV4amN2NGUza3VzWHlGSW4KY1d1Ym9qMWlWSzJHSTNhSW10NnZ6M05aUVRXNkZTazE2dEJEaDJEaUxqSGZjN0szcFRTdURkbGpOZHpCUmhRYQpoQlZpQ1Z2dkxEbER4Wm1LVlNld0QwbWkzb3RaSWF1Y1ZqVVFJOU1OKzJjNHRQTVhlTFJBMUx4dXZ4emF2WXIrClNIdU9xQzRabjV4R3J4dG9yeDk5c0pmMnRSVUJEL01DZ2dFQkFMRExRNzg0SDU2QjhlZU5zZE1wUk1mTitxZDQKclFlVTB2NUxUOElxdldTcmpLZTlHS3J4NTNSZ3RPNEtRTkdYc3FoMkROdktobFNDMys0amxCdms2UmxxUFdFYwpOSnRZSEw4UFBwby9hOEphM0F1RlpYL0NLNFFCeHR2SEdodjdPeG9JMG9zd0EzNG5NcVk2QXJaSktuSjEvcDljCmYxQk03TGRZQk1VNmVEeDBsSDVFM2xkM2lXVFN1ZUdWVk5PdzBpNmpoeDl3MUp0LzZwRis5NDJqdDFiRUoyN3YKYVdXT2REQ1g0SVIxMStiRlhhOEZJcEhCbStoTm1FdWRRc2hwN2pId2hCTjNiZnNSeHJXWGUyd1cvYkthdFBqWAo1N0p1bEFQVlN3L3h1TGJZZFZiVGlvdmRsMWxObXFJZEpqYVZma2ZZSzVJUVR1R0pxVHNzdVkvbWNITT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"
)

// Mock client
type MockHTTPClient struct {
	httpResponse *http.Response
	httpError    error
}

// Mock function
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.httpResponse, m.httpError
}

func TestCreateNewClient(t *testing.T) {
	t.Run("Success creation", func(t *testing.T) {
		mgnt, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		require.NoError(t, err)

		pk, err := loadPrivateKeyFromString(jwtPrivateKey)
		require.NoError(t, err)

		sm := jwtgo.GetSigningMethod(jwtAlgo)

		assert.Equal(t, err, nil)
		assert.Equal(t, mgnt.URL, URL)
		assert.Equal(t, mgnt.jwtIssuer, jwtIssuer)
		assert.Equal(t, mgnt.jwtSigningMethod, sm)
		assert.Equal(t, mgnt.jwtPrivateKey, pk)
	})

	t.Run("JWT issuer unset", func(t *testing.T) {
		_, err := New(URL, "", jwtAlgo, jwtPrivateKey)
		assert.EqualError(t, err, "JWT issuer unset")
	})

	t.Run("Invalid signing algorithm", func(t *testing.T) {
		_, err := New(URL, jwtIssuer, "RS999", jwtPrivateKey)
		assert.EqualError(t, err, "Unsupported signing method RS999")
	})

	t.Run("Invalid private key", func(t *testing.T) {
		_, err := New(URL, jwtIssuer, jwtAlgo, "")
		assert.EqualError(t, err, "Invalid Key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	})
}

func TestRequest(t *testing.T) {
	t.Run("Success request", func(t *testing.T) {
		mgnt, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		require.NoError(t, err)

		expected := `{"uid":"IDCA2AC08296","balance":"9916678.1751516791","locked":"280.0"}`
		body := ioutil.NopCloser(bytes.NewReader([]byte(expected)))
		mgnt.httpClient = &MockHTTPClient{
			httpResponse: &http.Response{
				StatusCode: 200,
				Body:       body,
			},
			httpError: nil,
		}

		res, apierr := mgnt.Request(http.MethodPost, "api/test", nil)

		assert.NotNil(t, res)
		assert.Nil(t, apierr)
		assert.Equal(t, res, []byte(expected))
	})

	t.Run("Invalid http method", func(t *testing.T) {
		mgnt, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		require.NoError(t, err)

		mgnt.httpClient = &MockHTTPClient{
			httpResponse: nil,
			httpError:    nil,
		}

		res, apierr := mgnt.Request(http.MethodGet, "api/test", nil)

		assert.Nil(t, res)
		assert.NotNil(t, apierr)
		assert.Equal(t, apierr.StatusCode, 500)
		assert.Equal(t, apierr.Error, "HTTP method is not allowed, accept only POST and PUT")
	})

	t.Run("HTTP client error", func(t *testing.T) {
		mgnt, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		require.NoError(t, err)

		mgnt.httpClient = &MockHTTPClient{
			httpResponse: nil,
			httpError:    fmt.Errorf("HTTP Error"),
		}

		res, apierr := mgnt.Request(http.MethodPost, "api/test", nil)

		assert.Nil(t, res)
		assert.NotNil(t, apierr)
		assert.Equal(t, apierr.StatusCode, 500)
		assert.Equal(t, apierr.Error, "HTTP Error")
	})

	t.Run("Invalid response with one error", func(t *testing.T) {
		mgnt, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		require.NoError(t, err)

		expected := `{"error":"Couldn't find record."}`
		body := ioutil.NopCloser(bytes.NewReader([]byte(expected)))
		mgnt.httpClient = &MockHTTPClient{
			httpResponse: &http.Response{
				StatusCode: 404,
				Body:       body,
			},
			httpError: nil,
		}

		res, apierr := mgnt.Request(http.MethodPost, "api/test", nil)

		assert.Nil(t, res)
		assert.NotNil(t, apierr)
		assert.Equal(t, apierr.StatusCode, 404)
		assert.Equal(t, apierr.Error, "Couldn't find record.")
		assert.Equal(t, len(apierr.Errors), 0)
	})

	t.Run("Invalid response with multiple errors", func(t *testing.T) {
		mgnt, err := New(URL, jwtIssuer, jwtAlgo, jwtPrivateKey)
		require.NoError(t, err)

		expected := `{"errors":["error_1","error_2"]}`
		body := ioutil.NopCloser(bytes.NewReader([]byte(expected)))
		mgnt.httpClient = &MockHTTPClient{
			httpResponse: &http.Response{
				StatusCode: 422,
				Body:       body,
			},
			httpError: nil,
		}

		res, apierr := mgnt.Request(http.MethodPost, "api/test", nil)

		assert.Nil(t, res)
		assert.NotNil(t, apierr)
		assert.Equal(t, apierr.StatusCode, 422)
		assert.Equal(t, apierr.Error, "")
		assert.Equal(t, len(apierr.Errors), 2)
		assert.Equal(t, apierr.Errors, []string{"error_1", "error_2"})
	})
}

func TestAllowedHTTPMethods(t *testing.T) {
	postMethod := allowedHTTPMethods(http.MethodPost)
	putMethod := allowedHTTPMethods(http.MethodPut)
	getMethod := allowedHTTPMethods(http.MethodGet)
	deleteMethod := allowedHTTPMethods(http.MethodDelete)
	unknownMethod := allowedHTTPMethods("xxx")
	emptyMethod := allowedHTTPMethods("")

	assert.Equal(t, postMethod, true)
	assert.Equal(t, putMethod, true)
	assert.Equal(t, getMethod, false)
	assert.Equal(t, deleteMethod, false)
	assert.Equal(t, unknownMethod, false)
	assert.Equal(t, emptyMethod, false)
}

func TestGenerateJWT(t *testing.T) {
	jwtPrivateKey := "LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlKS0FJQkFBS0NBZ0VBelNSeHpxZkhpVFg1bzl5N0JBdE1NM0lxcmtrSWNMZmhia3FHS0V3VFJXYkMyam5ZCmRDaXQ5SC9BV25zYTdpcnlOQ0hwS1lhUVozMkJ4MVVycnQrVk9kMm1YSDEwZHJ5VUtQcTZDdk1rSnBqYitNTncKTXlmd0dxKzdNdmMrUDBWcXp3dE5oNnplVThubVRzMTY2eWd0SVdzREEydWprc1R3ZHY3bEVFK0xMY1djbC90Ugp2dCtKcWtqeVNDYm1UOWl2bXUyeWh4UGFWbmU1TGxLQ2JnOWVJZEZTWEV1R2JFSnBpVGNhZ0lsSUh3VmJ6VnpSCllCbzlobjRXbHhXSmVUSkNQYmxIN0U0cmtrNUdJUDJqUnlvYmdUb2pTNERHS1hqc1ZwamJtc0l1Vk15OE5qQjUKWTJVcU54MVRaWFcvZTV2Nk9jUk9mOFFXSmxhQW1jNTd3S09La3Y0QXdoZ3QzWnluTXpnaVZaQ0Q2MzFDMmszYgpIUitTeHNneXFocHFZSEhUMThxM1hlVnI3d2lsR1Q1WnBremlIbk9SbjlFRjIxa21jczc2MWNBS3hIU3lENkNwCmdTVERObkc0Sm5sWjVnM1BpSk9YR1IvajgvZzNmZ3YydnAzL05nVm55Q2IrTUhCT0ZSaStXOXM1ZytMRU9XdTEKNmp1b2lWOHpaNlk5UmZFdHhVZjRzZ0ErNFJjUGZxQkJ5U2JPRW9neDJ5dDB3aVVJeUwyTEpZMmpBQUNxTVpSWApVUnB0bEtsVjAzbWZuL0I2aHkzNWR5VmN4ME9WRlpXK01tZUgzaHNHSWQwZE92UWdQaHpHOVBkcTVCUWVoQkp5ClpvdlU4RE54U0p5dnA1N0tqYlJWY2VIeklzUUw0MGZXdW81VlU0bHdnLzZueDdmYWlsdkRKSS9hZXQwQ0F3RUEKQVFLQ0FnQU1YRzdUSWY3L0FKYWJUaGlpeEwrQnRoWm1UQlpMSEhsaitPK2VpLzc1UnBqbEoya29qcTcwdGFIMAprY2hzbzMvV3JsaHJYU1ZrWndhajZUanBuNlZSU0U3VzhlUkxwMDlTTE5GN0NXMmJPY2kvYzU5V0pjanRBcnZICjlXZjF6Z3dDajg3TEp4cDZlQWI5cHBvS2czQTh2RU1CT01JeGZOWjBoU1Z1Vnl5dXhHS01NZU9hR2NRazA2SnQKd0pKT0syTmhkWU0xYW5mVWtBQkRqMHMyc0l4ZWcwdHdMa2phU3lJcTEzd3NWSmxZN1N5NzhpVFhvcDBrZG9LTAo5Z3REbDBpd2lYS1JCYURRZnhEd3VmZlZ1TzdSV1p4NDF6aVpsU1RBanhOa2Z1RGwwVFJpRzRlaytwcVJtWjNGCjFsT0Vja0NnckhpQ2NHRlpUQXNSdVlSeGRpbEtXSSthRU15L2lxUzFja0dPR0NuNUVMSmNpT21HV0hTQTF2eFgKZ1FMaVZ1Nkh1b0ZpcW1Jd3NWeHF5MFczN1ZDUEpVWFRXVmhCOGRXbjJscklyd3U0WWhhZzEwcXJ3cnJIQzFIWAo3a0hyU3JJRGdCSzBNdlQyTFh1TUU2eFcxT2k2N2k2RG05UHNrVXRIanVUNEpWTk94YUdRVUVBQm5YSkd5Y1MzCkRFVUFoR25qRmdpY29vM2JoQXZrdmh4WXNFdUdVRDZiVlZ1UTN0MytnbW0rWWVoSlUxdms4bldLYVdhOG14WEYKQWYrZmJIZ2c4TWxaVnJjOWpQQmZJZDQxdEZlUFpQbXpaaWFHRThnbHJ1d2xzQkhHa2F5WG4xYWtqM3JiV2NscgpLSWlRemJ4UjczTmZVTHNhZ2ZtL3lJWjBQeDRjWHgydG9zYUZiU3lkbGZLVDhaNGNEUUtDQVFFQTVyUVRwT3QzCjQ3VkZZUkw1ZTNaZTc3Rjc1MG9WN0pob3Q5YkZjenJFQTl6cVRFTC9Ua2pGbVE3RGR4aFgwTVo5YVZnMklMY3EKaThOdnpZWWhiMDlwd1ZsUWpzQmtocXA5WjlYWWU5bkNJRHBCMFhQWnRUM2pqbUtFcUYwZC8yT0FHeW16bWhVcwp6ajhnNDRsbFV5K1dvc1p2L2NicmF1RHZaUDBqbWd6cTRWSDlRK21VR2RKZlRDNXpBYldnZ3RodjhyWXd4YUVyCnhxOE9DNVp0V1BWK3BwN1dKV3ZSRWs4WWtTcFZVWjJVV2R0V2VaUk5hNWIxUHBUS3J4VGRxNkUvb3V4d1UwOTMKYWIwWDVlMGs3bnErVyttK1Y5Uks0NEUyL0x1c3JUWUlkRVRuWUpsRVo0aU9EUEJGNit5QlBtMVBKbGtSN3RjQwpNelhsUlNXRllUQjA0d0tDQVFFQTQ2TGNNUlZpdHd5QWkzamY1STd0amc4TmFpREw4dkJ1QzFzRXlMMndKaFNUClVXU0tiWmZIcEVpNlI1UVR2L01sd0RHNUZ4VThUWk1pSHl1RFFMR0pwYnQ3Vk5mZFVpWlozZkJiRkNKK3htaDMKOU1FRWszcWZJZlVHN0d3RmlkUEx4R3Q0OGMrREFNZW5nYnVhNlpMRTdvd21SZW8wK3k3cmRpQ3d3MEY3MWxUSApzbVd1aEhCa1hoTFlpU1lpZ1ZRTVJDR0U3aW5LdUdkQTdLd2dpTWQ4VVhrV3NibDJZRmFmd3JSRlVZVmJRUlJ2ClZTVnVMYVVoYTZCbzFkbGxLZ1lrVFIwOEZRNWhrVisvaGtZNlZoM1BCRzZEZXhJR1FiNlJrQjdDNWprSVhDdlMKUzFvOWF6aENreTFralo4YXpuMUg2Ky9xM25qckl5UXNtcWdoZUpBZFB3S0NBUUJBNG1Tai9aVzZkVUVPREVnZQpjU3hDUGFpYlpEckdVQmNqblVQckpKdjhlaVZyVFd5QWwvYjdGU3ZrVXZSZnczT0NMVTBMNW5nUTF1YWE1eDZBCkw5V09pNUFjbGYrdjRFTms4TC95RlV5RHc5Ni9DZFl4SXpiYzFOaDZnYlh1SGczcGxkRHRoUWNVK3F4RlVsOHQKQmpWWGtuZnM2QVZPQ2ZWS2NlZVJiQkNqVG12c3JjVDVmakZQTzhFY3VmaHExSFNuenBYby8ydFFkZXQ5VnRGcQpNNkZyTzBEL1JWT0gwcmNXSE5IaUltK1cxaGw4R0RtdUNNYncwdWd1VmJBQ2xWZFFleThjUHoxV2Y5ZzQwbm1RCm1QVHc1TXlqNXhFbzZ5Nkw1anlxZW9mbUszcm5zRE9NNnRzSXlJcmh6NktKN0RSV2xMWjJkZ0lvWlFBV2NuY1EKM3BBQkFvSUJBQmQ2Q1dtS2doYk0xRWtPRzFFd0tISFpQWkh2ZGZsRk1LUTlLOTRrS2hHVFY2b3lTMUNJTWMvUQpyRjJMZVFuMzRySFNydnNoZG9tdG5meEcrWTluZ0FHMnR6NkYwTTZUSS91T3VXWDNOTW56cGtONDBLY0JJMzVXCkRmTytKRWdWcnROQUhrWWFGN0d4NWFXc21vcHlWNXNlbXlma3dyZ1JHN21nSDNyVHV4amN2NGUza3VzWHlGSW4KY1d1Ym9qMWlWSzJHSTNhSW10NnZ6M05aUVRXNkZTazE2dEJEaDJEaUxqSGZjN0szcFRTdURkbGpOZHpCUmhRYQpoQlZpQ1Z2dkxEbER4Wm1LVlNld0QwbWkzb3RaSWF1Y1ZqVVFJOU1OKzJjNHRQTVhlTFJBMUx4dXZ4emF2WXIrClNIdU9xQzRabjV4R3J4dG9yeDk5c0pmMnRSVUJEL01DZ2dFQkFMRExRNzg0SDU2QjhlZU5zZE1wUk1mTitxZDQKclFlVTB2NUxUOElxdldTcmpLZTlHS3J4NTNSZ3RPNEtRTkdYc3FoMkROdktobFNDMys0amxCdms2UmxxUFdFYwpOSnRZSEw4UFBwby9hOEphM0F1RlpYL0NLNFFCeHR2SEdodjdPeG9JMG9zd0EzNG5NcVk2QXJaSktuSjEvcDljCmYxQk03TGRZQk1VNmVEeDBsSDVFM2xkM2lXVFN1ZUdWVk5PdzBpNmpoeDl3MUp0LzZwRis5NDJqdDFiRUoyN3YKYVdXT2REQ1g0SVIxMStiRlhhOEZJcEhCbStoTm1FdWRRc2hwN2pId2hCTjNiZnNSeHJXWGUyd1cvYkthdFBqWAo1N0p1bEFQVlN3L3h1TGJZZFZiVGlvdmRsMWxObXFJZEpqYVZma2ZZSzVJUVR1R0pxVHNzdVkvbWNITT0KLS0tLS1FTkQgUlNBIFBSSVZBVEUgS0VZLS0tLS0K"
	jwtIssuer := "rice.com"
	payload := map[string]interface{}{"bar": "baz"}

	mgnt, err := New("", jwtIssuer, "", jwtPrivateKey)
	require.NoError(t, err)

	jwt, err := mgnt.generateJWT(payload, time.Hour, int64(1611833128), "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9")
	require.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"payload": "eyJkYXRhIjp7ImJhciI6ImJheiJ9LCJleHAiOjE2MTE4MzY3MjgsImlhdCI6MTYxMTgzMzEyOCwiaXNzIjoicmljZS5jb20iLCJqdGkiOiJleUpoYkdjaU9pSlNVekkxTmlJc0luUjVjQ0k2SWtwWFZDSjkifQ",
		"signatures": []map[string]interface{}{
			{
				"header":    map[string]string{"kid": "rice.com"},
				"protected": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9",
				"signature": "wCLJhZ5I0d_mN_lP7tT_uy80qr2kaAzf--EKpQdMhbRDhQDNEl7XDMoBssqMPT6YF7gpNk_7npijZA_YrBYBib5EZZDm3fte5l88xK9r4Xwg_eKxq6YuXrJS3XtxK_Ybh93PeAKThUlKxTyw71AEDmiy7SmWgEpqbWX6hz4I4zUFQTu0U981Ava8l5BNf1Nz76kIE5jKBYZbO2nn96TVHgZPHhUZ1SPm3qNwpZuKFwnjI22-m3n7M18D5cg-6IYxbyXahnrIyy3wlbPK4UxkVfZwFwDnpGfYxCvS5SIIW_0fR7OGVA_PEDpJgswClPioyh2rV6wztsqBubxEk1H1sEUWp_DZPWCBbvxMr1pYubk-AwYHpR9patSR-eZugpXkfXCHFy06Tqrnlt18eXZmpbGYHyYCj4KGiHsliigmV_-ssl6bgeaCbr4Hde7nFb94XvF8a4UQN9KyQImKcW_bCqmWEvDg3opK3BLfBMdQqJIQHTUqwPhZ61EHvcLZvgiQCYk0Yq96IO8qNZPoJAKyvsNVV0iak6vqYs1IzODGtQal-Nn5JWowo0pK2sHpcETPcidnB_TsxLZsK1XUtkTXNqzYa424V1RJb3y94ylNLLPpn9-t4ouk_5DT2iDYRf6Qod7l2fyJNgUaGYjzo7DdrHN1Kb9UTUxZ1g4yNe44HBc"},
		},
	}, jwt)
}
