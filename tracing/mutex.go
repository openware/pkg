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
		logger.Trace("waiting for lock", "session", session, "queueLen", waiting)
	}
	sm.waiters = append(sm.waiters, session)
	sm.waitersLock.Unlock()

	sm.realLock.Lock()

	sm.waitersLock.Lock()
	behindLen := len(sm.waiters) - 1
	sm.waitersLock.Unlock()

	logger.Trace("acquired lock", "session", session, "queueLen", behindLen)
}

// Unlock releases the lock.
func (sm *Mutex) Unlock() {
	sm.waitersLock.Lock()
	// remove self from the list of waiters
	if len(sm.waiters) > 0 {
		sm.waiters = sm.waiters[1:]
	}
	sm.waitersLock.Unlock()

	sm.realLock.Unlock()
}
