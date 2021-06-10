package exchangerates

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/openware/pkg/currency/forexprovider/base"
)

var e ExchangeRates

var initialSetup bool
type mockRoundTripper struct {
	mock.Mock
}

func (m *mockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	args := m.Called(request)
	return args.Get(0).(*http.Response), args.Error(1)
}

func setup() *ExchangeRates {
	var e ExchangeRates
	e.Setup(base.Settings{
		Name:    "ExchangeRates",
		Enabled: true,
		APIKey: "something",
		Verbose: true,
	})
	return &e
}

func seedRate(date time.Time, base string, rates map[string]float64) map[string] interface{} {
	return map[string]interface{}{
		"timestamp": date.Unix(),
		"base": base,
		"date": date.Format("2006-01-02"),
		"rates": rates,
	}
}
func mockRate(base string, rates map[string]float64) io.ReadCloser {
	b, _ := json.Marshal(seedRate(time.Now(), base, rates))
	return ioutil.NopCloser(bytes.NewReader(b))
}

func mockHistoryRate(date time.Time, base string, rates map[string]float64) io.ReadCloser {
	seed := seedRate(date, base, rates)
	seed["historical"] = true
	b, _ := json.Marshal(seed)

	return ioutil.NopCloser(bytes.NewReader(b))
}

func mockTimeseries(start, end time.Time, base string, rates ...map[string]float64) io.ReadCloser {
	seed := seedRate(start, base, nil)
	delete(seed, "rates")

	diff := int(math.Ceil(end.Sub(start).Hours() / 24))
	dateFormat := "2006-01-02"
	rateResponse := map[string]map[string]float64{}

	for i := 0; i < diff; i ++ {
		rateResponse[start.Add(time.Hour * time.Duration(i * 24)).Format(dateFormat)] = rates[i]
	}
	seed["rates"] = rateResponse

	b, _ := json.Marshal(seed)

	return ioutil.NopCloser(bytes.NewReader(b))
}

func TestGetLatestRates(t *testing.T) {
	t.Parallel()

	e := setup()
	m := &mockRoundTripper{}

	m.
		On("RoundTrip", mock.Anything).Once().Return(&http.Response{
		Status:           "ok",
		StatusCode:       200,
		Body:             mockRate("USD", map[string]float64{"EUR": 0.950345, "USD": 1.00}),
	}, nil)
	m.On("RoundTrip", mock.Anything).Once().Return(&http.Response{
		Status:           "ok",
		StatusCode:       200,
		Body:             mockRate("EUR", map[string]float64{"EUR": 1, "AUD": 1.912912}),
	}, nil)

	e.Requester.HTTPClient.Transport = m

	result, err := e.GetLatestRates("USD", "")

	m.AssertNumberOfCalls(t, "RoundTrip", 1)
	req := m.Calls[0].Arguments[0].(*http.Request)

	require.Nil(t, err)
	//require.Equal(t, "USD", result.Base)
	require.Equal(t, "USD", req.URL.Query().Get("base"))
	require.Equal(t, "", req.URL.Query().Get("symbols"))
	require.GreaterOrEqual(t, len(result.Rates), 1)

	result2, err := e.GetLatestRates("USD", "")

	m.AssertNumberOfCalls(t, "RoundTrip", 2)

	req = m.Calls[1].Arguments[0].(*http.Request)

	require.Nil(t, err)
	//require.Equal(t, "USD", result.Base)
	require.Equal(t, "USD", req.URL.Query().Get("base"))
	require.Equal(t, "", req.URL.Query().Get("symbols"))
	require.GreaterOrEqual(t, len(result2.Rates), 1)
}

func TestCleanCurrencies(t *testing.T) {
	t.Parallel()

	result := cleanCurrencies("USD", "USD,AUD")
	if result != "AUD" {
		t.Fatalf("unexpected result. AUD should be the only symbol")
	}

	result = cleanCurrencies("", "EUR,USD")
	if result != "USD" {
		t.Fatalf("unexpected result. USD should be the only symbol")
	}

	if cleanCurrencies("EUR", "RUR") != "RUB" {
		t.Fatalf("unexpected result. RUB should be the only symbol")
	}

	if cleanCurrencies("EUR", "AUD,BLA") != "AUD" {
		t.Fatalf("unexpected result. AUD should be the only symbol")
	}
}

func TestGetFailedRates(t *testing.T) {
	t.Parallel()

	e := setup()
	originalRoundtripper := e.Requester.HTTPClient.Transport
	m := &mockRoundTripper{}

	m.
		On("RoundTrip", mock.Anything).Once().Return(&http.Response{
		Status:           "fail",
		StatusCode:       500,
		Body:             ioutil.NopCloser(strings.NewReader(`{"error":"internal server error"}`)),
	}, nil)

	e.Requester.HTTPClient.Transport = m

	_, err := e.GetLatestRates("USD", "AUD")
	require.NotNil(t, err)
	m.AssertNumberOfCalls(t, "RoundTrip", 1)

	req, _ := m.Calls[0].Arguments[0].(*http.Request)

	require.Equal(t, "USD", req.URL.Query().Get("base"))
	require.Equal(t, "AUD", req.URL.Query().Get("symbols"))

	e.Requester.HTTPClient.Transport = originalRoundtripper
}

func TestGetHistoricalRates(t *testing.T) {
	t.Parallel()

	e := setup()
	m := &mockRoundTripper{}
	expectedDate := "2010-01-12"
	date, _ := time.Parse("2006-01-02", expectedDate)
	m.On("RoundTrip", mock.Anything).Once().Return(&http.Response{
		StatusCode: 200,
		Body: mockHistoryRate(date, "USD", map[string]float64{"AUD": 1.2323, "EUR": 0.9999}),
	}, nil)

	e.Requester.HTTPClient.Transport = m
	_, err := e.GetHistoricalRates("-1", "USD", []string{"AUD"})
	require.NotNil(t, err)
	m.AssertNumberOfCalls(t, "RoundTrip", 0)

	_, err = e.GetHistoricalRates(expectedDate, "USD", []string{"EUR,AUD"})
	require.Nil(t, err)
	m.AssertNumberOfCalls(t, "RoundTrip", 1)

	req := m.Calls[0].Arguments[0].(*http.Request)

	require.Equal(t, "USD", req.URL.Query().Get("base"))
	require.Contains(t, req.URL.Path, expectedDate)
	require.Equal(t, "EUR,AUD", req.URL.Query().Get("symbols"))
}

func TestGetTimeSeriesRates(t *testing.T) {
	t.Parallel()

	e := setup()
	m := &mockRoundTripper{}
	expectedDate := "2010-01-12"
	date, _ := time.Parse(dateLayout, expectedDate)
	diff := 24 * 3
	duration := time.Hour * time.Duration(diff)
	sub := date.Add(duration)
	m.On("RoundTrip", mock.Anything).Once().Return(&http.Response{
		StatusCode: 200,
		Body: mockTimeseries(
			date, sub, "USD",
			map[string]float64{"AUD": 1.2323, "EUR": 0.9999},
			map[string]float64{"AUD": 1.2323, "EUR": 0.9999},
			map[string]float64{"AUD": 1.2323, "EUR": 0.9999},
		),
	}, nil)

	e.Requester.HTTPClient.Transport = m
	_, err := e.GetTimeSeriesRates("", "", "USD", []string{"EUR", "USD"})
	require.NotNil(t, err)
	// 0 means no http calls should've been made due to validation error
	m.AssertNumberOfCalls(t, "RoundTrip", 0)

	_, err = e.GetTimeSeriesRates("-1", "-1", "USD", []string{"EUR,USD"})
	require.NotNil(t, err)
	// 0 means no http calls should've been made due to validation error
	m.AssertNumberOfCalls(t, "RoundTrip", 0)

	_, err = e.GetTimeSeriesRates(date.Format(dateLayout), sub.Format(dateLayout), "USD", []string{"EUR,AUD"})

	m.AssertNumberOfCalls(t, "RoundTrip", 1)

	req, _ := m.Calls[0].Arguments[0].(*http.Request)

	require.Equal(t, "USD", req.URL.Query().Get("base"))
	require.Equal(t, expectedDate, req.URL.Query().Get("start_date"))
	require.Equal(t, "2010-01-15", req.URL.Query().Get("end_date"))
	require.Contains(t, req.URL.Path, "timeseries")
	require.Equal(t, "EUR,AUD", req.URL.Query().Get("symbols"))
}
