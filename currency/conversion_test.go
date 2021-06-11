package currency

import (
	"fmt"
	exchangerates "github.com/openware/pkg/currency/forexprovider/exchangeratesapi.io"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

type mockRoundTripper struct {
	mock.Mock
}

func (m *mockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	args := m.Called(request)
	return args.Get(0).(*http.Response), args.Error(1)
}
func setupMock()  {
	mockRoundTrip := &mockRoundTripper{}

	mockBodyString := `{
    "success": true,
    "timestamp": 1519296206,
    "base": "AUD",
    "date": "2021-03-17",
    "rates": {
        "EUR": 0.566015,
        "AUD": 1,
        "CAD": 1.560132,
        "CHF": 1.154727,
        "CNY": 7.827874,
        "GBP": 0.882047,
        "JPY": 132.360679,
        "USD": 1.23396
    }
}`
	mockBody := ioutil.NopCloser(strings.NewReader(mockBodyString))

	mockRoundTrip.On("RoundTrip", mock.Anything).Once().Return(&http.Response{
		StatusCode: 200,
		Status: "OK",
		Body: mockBody,
	}, nil)

	storage.fiatExchangeMarkets.Primary.Provider.(*exchangerates.ExchangeRates).Requester.HTTPClient.Transport = mockRoundTrip
}

func TestNewConversionFromString(t *testing.T) {
	setupMock()
	expected := "AUDUSD"
	conv, err := NewConversionFromString(expected)
	require.Nil(t, err)
	require.Equal(t, expected, conv.String())

	newexpected := strings.ToLower(expected)
	conv, err = NewConversionFromString(newexpected)
	require.Nil(t, err)
	require.Equal(t, newexpected, conv.String())
}

func TestNewConversionFromStrings(t *testing.T) {
	setupMock()
	from := "AUD"
	to := "USD"
	expected := "AUDUSD"

	conv, err := NewConversionFromStrings(from, to)
	if err != nil {
		t.Error(err)
	}

	if conv.String() != expected {
		t.Errorf("NewConversion() error expected %s but received %s",
			expected,
			conv)
	}
}

func TestNewConversion(t *testing.T) {
	setupMock()
	from := NewCode("AUD")
	to := NewCode("USD")
	expected := "AUDUSD"

	conv, err := NewConversion(from, to)
	if err != nil {
		t.Error(err)
	}

	if conv.String() != expected {
		t.Errorf("NewConversion() error expected %s but received %s",
			expected,
			conv)
	}
}

func TestConversionIsInvalid(t *testing.T) {
	setupMock()
	from := AUD
	to := USD

	conv, err := NewConversion(from, to)
	if err != nil {
		t.Fatal(err)
	}

	if conv.IsInvalid() {
		t.Errorf("IsInvalid() error expected false but received %v",
			conv.IsInvalid())
	}

	to = AUD
	conv, err = NewConversion(from, to)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestConversionIsFiatPair(t *testing.T) {
	setupMock()
	from := AUD
	to := USD

	conv, err := NewConversion(from, to)
	if err != nil {
		t.Fatal(err)
	}

	if !conv.IsFiat() {
		t.Errorf("IsFiatPair() error expected true but received %v",
			conv.IsFiat())
	}

	to = LTC
	conv, err = NewConversion(from, to)
	if err == nil {
		t.Error("Expected error")
	}
}

func TestConversionsRatesSystem(t *testing.T) {
	var SuperDuperConversionSystem ConversionRates

	if SuperDuperConversionSystem.HasData() {
		t.Fatalf("HasData() error expected false but received %v",
			SuperDuperConversionSystem.HasData())
	}

	testmap := map[string]float64{
		"USDAUD": 1.3969317581,
		"USDBRL": 3.7047257979,
		"USDCAD": 1.3186386881,
		"USDCHF": 1,
		"USDCNY": 6.7222712044,
		"USDCZK": 22.6406277552,
		"USDDKK": 6.5785575736,
		"USDEUR": 0.8816787163,
		"USDGBP": 0.7665755599,
		"USDHKD": 7.8492329395,
		"USDILS": 3.6152354082,
		"USDINR": 71.154558279,
		"USDJPY": 110.7476635514,
		"USDKRW": 1122.7913948157,
		"USDMXN": 19.1589666725,
		"USDNOK": 8.5818197849,
		"USDNZD": 1.4559160642,
		"USDPLN": 3.8304531829,
		"USDRUB": 65.7533062952,
		"USDSEK": 9.3196085346,
		"USDSGD": 1.3512608006,
		"USDTHB": 31.0950449656,
		"USDZAR": 14.138070887,
	}

	err := SuperDuperConversionSystem.Update(testmap)
	if err != nil {
		t.Fatal(err)
	}

	err = SuperDuperConversionSystem.Update(nil)
	if err == nil {
		t.Fatal("Update() error cannot be nil")
	}

	if !SuperDuperConversionSystem.HasData() {
		t.Fatalf("HasData() error expected true but received %v",
			SuperDuperConversionSystem.HasData())
	}

	// * to a rate
	p := SuperDuperConversionSystem.m[USD.Item][AUD.Item]
	// inverse * to a rate
	pi := SuperDuperConversionSystem.m[AUD.Item][USD.Item]
	r := *p * 1000
	expectedRate := 1396.9317581
	if r != expectedRate {
		t.Errorf("Convert() error expected %.13f but received %.13f",
			expectedRate,
			r)
	}

	inverseR := *pi * expectedRate
	expectedInverseRate := float64(1000)
	if inverseR != expectedInverseRate {
		t.Errorf("Convert() error expected %.13f but received %.13f",
			expectedInverseRate,
			inverseR)
	}
}

func TestGetRate(t *testing.T) {
	setupMock()
	from := NewCode("AUD")
	to := NewCode("USD")

	c, err := NewConversion(from, to)
	if err != nil {
		t.Error(err)
	}
	rate, err := c.GetRate()
	if err != nil {
		t.Error(err)
	}
	if rate == 0 {
		t.Error("Rate not set")
	}
	inv, err := c.GetInversionRate()
	if err != nil {
		t.Error(err)
	}
	if inv == 0 {
		t.Error("Inverted rate not set")
	}
	conv, err := c.Convert(1)
	if err != nil {
		t.Error(err)
	}
	if rate != conv {
		t.Errorf("Incorrect rate %v %v", rate, conv)
	}
	invConv, err := c.ConvertInverse(1)
	if err != nil {
		t.Error(err)
	}
	if inv != invConv {
		t.Errorf("Incorrect rate %v %v", conv, invConv)
	}

	var convs ConversionRates
	var convRate float64
	_, err = convs.GetRate(BTC, USDT)
	if err == nil {
		t.Errorf("Expected %s", fmt.Errorf("rate not found for from %s to %s conversion",
			BTC,
			USD))
	}
	convRate, err = convs.GetRate(USDT, USD)
	if err != nil {
		t.Error(err)
	}
	if convRate != 1 {
		t.Errorf("Expected rate to be 1")
	}

	convRate, err = convs.GetRate(RUR, RUB)
	if err != nil {
		t.Error(err)
	}
	if convRate != 1 {
		t.Errorf("Expected rate to be 1")
	}

	convRate, err = convs.GetRate(RUB, RUR)
	if err != nil {
		t.Error(err)
	}
	if convRate != 1 {
		t.Errorf("Expected rate to be 1")
	}
}
