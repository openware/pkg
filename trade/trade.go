package trade

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gofrs/uuid"
	"github.com/openware/pkg/common"
	"github.com/openware/pkg/kline"
	"github.com/openware/pkg/log"
	"github.com/openware/pkg/order"
)

// Setup creates the trade processor if trading is supported
func (p *Processor) setup(wg *sync.WaitGroup) {
	p.mutex.Lock()
	p.bufferProcessorInterval = BufferProcessorIntervalTime
	p.mutex.Unlock()
	go p.Run(wg)
}

// AddTradesToBuffer will push trade data onto the buffer
func AddTradesToBuffer(exchangeName string, data ...Data) error {
	if len(data) == 0 {
		return nil
	}
	var errs common.Errors
	if atomic.AddInt32(&processor.started, 0) == 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		processor.setup(&wg)
		wg.Wait()
	}
	var validDatas []Data
	for i := range data {
		if data[i].Price == 0 ||
			data[i].Amount == 0 ||
			data[i].CurrencyPair.IsEmpty() ||
			data[i].Exchange == "" ||
			data[i].Timestamp.IsZero() {
			errs = append(errs, fmt.Errorf("%v received invalid trade data: %+v", exchangeName, data[i]))
			continue
		}

		if data[i].Price < 0 {
			data[i].Price *= -1
			data[i].Side = order.Sell
		}
		if data[i].Amount < 0 {
			data[i].Amount *= -1
			data[i].Side = order.Sell
		}
		if data[i].Side == order.Bid {
			data[i].Side = order.Buy
		}
		if data[i].Side == order.Ask {
			data[i].Side = order.Sell
		}
		uu, err := uuid.NewV4()
		if err != nil {
			errs = append(errs, fmt.Errorf("%s uuid failed to generate for trade: %+v", exchangeName, data[i]))
		}
		data[i].ID = uu
		validDatas = append(validDatas, data[i])
	}
	processor.mutex.Lock()
	processor.buffer = append(processor.buffer, validDatas...)
	processor.mutex.Unlock()
	if len(errs) > 0 {
		return errs
	}
	return nil
}

// Run will save trade data to the database in batches
func (p *Processor) Run(wg *sync.WaitGroup) {
	wg.Done()
	if !atomic.CompareAndSwapInt32(&p.started, 0, 1) {
		log.Error(log.Trade, "trade processor already started")
		return
	}
	defer func() {
		atomic.CompareAndSwapInt32(&p.started, 1, 0)
	}()
	p.mutex.Lock()
	ticker := time.NewTicker(p.bufferProcessorInterval)
	p.mutex.Unlock()
	for {
		<-ticker.C
		p.mutex.Lock()
		bufferCopy := append(p.buffer[:0:0], p.buffer...)
		p.buffer = nil
		p.mutex.Unlock()
		if len(bufferCopy) == 0 {
			ticker.Stop()
			return
		}
	}
}

// ConvertTradesToCandles turns trade data into kline.Items
func ConvertTradesToCandles(interval kline.Interval, trades ...Data) (kline.Item, error) {
	if len(trades) == 0 {
		return kline.Item{}, errors.New("no trades supplied")
	}
	groupedData := groupTradesToInterval(interval, trades...)
	candles := kline.Item{
		Exchange: trades[0].Exchange,
		Pair:     trades[0].CurrencyPair,
		Asset:    trades[0].AssetType,
		Interval: interval,
	}
	for k, v := range groupedData {
		candles.Candles = append(candles.Candles, classifyOHLCV(time.Unix(k, 0), v...))
	}

	return candles, nil
}

func groupTradesToInterval(interval kline.Interval, times ...Data) map[int64][]Data {
	groupedData := make(map[int64][]Data)
	for i := range times {
		nearestInterval := getNearestInterval(times[i].Timestamp, interval)
		groupedData[nearestInterval] = append(
			groupedData[nearestInterval],
			times[i],
		)
	}
	return groupedData
}

func getNearestInterval(t time.Time, interval kline.Interval) int64 {
	return t.Truncate(interval.Duration()).UTC().Unix()
}

func classifyOHLCV(t time.Time, datas ...Data) (c kline.Candle) {
	sort.Sort(ByDate(datas))
	c.Open = datas[0].Price
	c.Close = datas[len(datas)-1].Price
	for i := range datas {
		if datas[i].Price < 0 {
			datas[i].Price *= -1
		}
		if datas[i].Amount < 0 {
			datas[i].Amount *= -1
		}
		if datas[i].Price < c.Low || c.Low == 0 {
			c.Low = datas[i].Price
		}
		if datas[i].Price > c.High {
			c.High = datas[i].Price
		}
		c.Volume += datas[i].Amount
	}
	c.Time = t
	return c
}

// FilterTradesByTime removes any trades that are not between the start
// and end times
func FilterTradesByTime(trades []Data, startTime, endTime time.Time) []Data {
	if startTime.IsZero() || endTime.IsZero() {
		// can't filter without boundaries
		return trades
	}
	var filteredTrades []Data
	for i := range trades {
		if trades[i].Timestamp.After(startTime) && trades[i].Timestamp.Before(endTime) {
			filteredTrades = append(filteredTrades, trades[i])
		}
	}

	return filteredTrades
}
