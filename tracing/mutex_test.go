package tracing

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/openware/pkg/log"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/baggage"
)

func TestMutex(t *testing.T) {
	logger := log.NewLogger("test")
	sessionBaggageKey := "request"
	lock := NewMutex(sessionBaggageKey, logger)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		bggMember, err := baggage.NewMember(sessionBaggageKey, fmt.Sprint(i))
		require.NoError(t, err)

		bgg, err := baggage.New(bggMember)
		require.NoError(t, err)

		wg.Add(1)
		go func() {
			lock.Lock(baggage.ContextWithBaggage(context.Background(), bgg))

			time.Sleep(100 * time.Millisecond)

			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
}
