package core

import "sync"

var(
	defaultTraceIDKey="trace_id"
	l =new(sync.RWMutex)
)

func setTraceIDKey(traceIDKey string)  {

	defer l.Unlock()
	l.Lock()
	defaultTraceIDKey=traceIDKey

}

func getTraceIDKey()string  {

	defer l.RUnlock()
	l.RLock()
	return defaultTraceIDKey

}
