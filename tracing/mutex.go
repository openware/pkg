package tracing

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/openware/pkg/log"
)

// Mutex is just like the standard sync.Mutex, except that it is aware of the Context
// and logs some diagnostic information into the current span.
type Mutex struct {
	sessionBaggageKey string
	logger            log.CtxLogger

	realLock sync.Mutex

	waiters     []string
	waitersLock sync.Mutex
}

// NewMutex creates a new mutex that logs diagnostic information into the current span.
func NewMutex(sessionBaggageKey string, logger log.CtxLogger) *Mutex {
	return &Mutex{
		sessionBaggageKey: sessionBaggageKey,
		logger:            logger,
	}
}

// Lock acquires an exclusive lock.
func (sm *Mutex) Lock(ctx context.Context) {
	logger := sm.logger.For(ctx)
	session := BaggageItem(ctx, sm.sessionBaggageKey)
	activeSpan := trace.SpanFromContext(ctx)
	activeSpan.SetAttributes(attribute.String(sm.sessionBaggageKey, session))

	sm.waitersLock.Lock()
	if waiting := len(sm.waiters); waiting > 0 {
		logger.Info("waiting for lock", "session", session, "waiting", waiting, "waiters", sm.waitersString(0, waiting))
	}
	sm.waiters = append(sm.waiters, session)
	sm.waitersLock.Unlock()

	sm.realLock.Lock()

	sm.waitersLock.Lock()
	behindLen := len(sm.waiters) - 1
	sm.waitersLock.Unlock()

	logger.Info("acquired lock", "session", session, "waiting", behindLen, "waiters", sm.waitersString(0, behindLen+1))
}

// Unlock releases the lock.
func (sm *Mutex) Unlock() {
	sm.waitersLock.Lock()
	// remove self from the list of waiters
	if len(sm.waiters) > 0 {
		sm.logger.Info("released lock", "session", sm.waiters[0], "waiting", len(sm.waiters), "waiters", sm.waitersString(1, len(sm.waiters)))
		sm.waiters = sm.waiters[1:]
	}
	sm.waitersLock.Unlock()

	sm.realLock.Unlock()
}

func (sm *Mutex) waitersString(from, to int) string {
	if from >= to || from < 0 || from >= len(sm.waiters) {
		return ""
	}

	if to > len(sm.waiters) {
		to = len(sm.waiters)
	}

	str := sm.waiters[from]
	for i := from + 1; i < to; i++ {
		str += "," + sm.waiters[i]
	}

	return str
}
