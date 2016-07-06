// This file was generated by counterfeiter
package middlewarefakes

import (
	"net/http"
	"sync"

	"github.com/cloudfoundry-incubator/switchboard/api/middleware"
)

type FakeMiddleware struct {
	WrapStub        func(http.Handler) http.Handler
	wrapMutex       sync.RWMutex
	wrapArgsForCall []struct {
		arg1 http.Handler
	}
	wrapReturns struct {
		result1 http.Handler
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMiddleware) Wrap(arg1 http.Handler) http.Handler {
	fake.wrapMutex.Lock()
	fake.wrapArgsForCall = append(fake.wrapArgsForCall, struct {
		arg1 http.Handler
	}{arg1})
	fake.recordInvocation("Wrap", []interface{}{arg1})
	fake.wrapMutex.Unlock()
	if fake.WrapStub != nil {
		return fake.WrapStub(arg1)
	} else {
		return fake.wrapReturns.result1
	}
}

func (fake *FakeMiddleware) WrapCallCount() int {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return len(fake.wrapArgsForCall)
}

func (fake *FakeMiddleware) WrapArgsForCall(i int) http.Handler {
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return fake.wrapArgsForCall[i].arg1
}

func (fake *FakeMiddleware) WrapReturns(result1 http.Handler) {
	fake.WrapStub = nil
	fake.wrapReturns = struct {
		result1 http.Handler
	}{result1}
}

func (fake *FakeMiddleware) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.wrapMutex.RLock()
	defer fake.wrapMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMiddleware) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ middleware.Middleware = new(FakeMiddleware)