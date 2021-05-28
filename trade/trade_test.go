package trade

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/openware/pkg/asset"
	"github.com/openware/pkg/currency"
	"github.com/openware/pkg/kline"
	"github.com/openware/pkg/order"
)

func TestConvertTradesToCandles(t *testing.T) {
	t.Parallel()
	cp, _ := currency.NewPairFromString("BTC-USD")
	startDate := time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC)
	candles, err := ConvertTradesToCandles(kline.FifteenSecond, []Data{
		{
			Timestamp:    startDate,
			Exchange:     "test!",
			CurrencyPair: cp,
			AssetType:    asset.Spot,
			Price:        1337,
			Amount:       1337,
			Side:         order.Buy,
		},
		{
			Timestamp:    startDate.Add(time.Second),
			Exchange:     "test!",
			CurrencyPair: cp,
			AssetType:    asset.Spot,
			Price:        1337,
			Amount:       1337,
			Side:         order.Buy,
		},
		{
			Timestamp:    startDate.Add(time.Minute),
			Exchange:     "test!",
			CurrencyPair: cp,
			AssetType:    asset.Spot,
			Price:        -1337,
			Amount:       -1337,
			Side:         order.Buy,
		},
	}...)
	if err != nil {
		t.Fatal(err)
	}
	if len(candles.Candles) != 2 {
		t.Fatal("unexpected candle amount")
	}
	if candles.Interval != kline.FifteenSecond {
		t.Error("expected fifteen seconds")
	}
}

func TestShutdown(t *testing.T) {
	t.Parallel()
	var p Processor
	p.mutex.Lock()
	p.bufferProcessorInterval = time.Second
	p.mutex.Unlock()
	var wg sync.WaitGroup
	wg.Add(1)
	go p.Run(&wg)
	wg.Wait()
	time.Sleep(time.Millisecond)
	if atomic.LoadInt32(&p.started) != 1 {
		t.Error("expected it to start running")
	}
	time.Sleep(time.Second * 2)
	if atomic.LoadInt32(&p.started) != 0 {
		t.Error("expected it to stop running")
	}
}

func TestFilterTradesByTime(t *testing.T) {
	t.Parallel()
	trades := []Data{
		{
			Exchange:  "test",
			Timestamp: time.Now().Add(-time.Second),
		},
	}
	trades = FilterTradesByTime(trades, time.Now().Add(-time.Minute), time.Now())
	if len(trades) != 1 {
		t.Error("failed to filter")
	}
	trades = FilterTradesByTime(trades, time.Now().Add(-time.Millisecond), time.Now())
	if len(trades) != 0 {
		t.Error("failed to filter")
	}
}
