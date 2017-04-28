// This file was generated by counterfeiter
package auth0fakes

import (
	"sync"

	"github.com/3dsim/auth0/auth0"
)

type FakeTokenFetcher struct {
	TokenStub        func(audience string) (string, error)
	tokenMutex       sync.RWMutex
	tokenArgsForCall []struct {
		audience string
	}
	tokenReturns struct {
		result1 string
		result2 error
	}
	tokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTokenFetcher) Token(audience string) (string, error) {
	fake.tokenMutex.Lock()
	ret, specificReturn := fake.tokenReturnsOnCall[len(fake.tokenArgsForCall)]
	fake.tokenArgsForCall = append(fake.tokenArgsForCall, struct {
		audience string
	}{audience})
	fake.recordInvocation("Token", []interface{}{audience})
	fake.tokenMutex.Unlock()
	if fake.TokenStub != nil {
		return fake.TokenStub(audience)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.tokenReturns.result1, fake.tokenReturns.result2
}

func (fake *FakeTokenFetcher) TokenCallCount() int {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return len(fake.tokenArgsForCall)
}

func (fake *FakeTokenFetcher) TokenArgsForCall(i int) string {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return fake.tokenArgsForCall[i].audience
}

func (fake *FakeTokenFetcher) TokenReturns(result1 string, result2 error) {
	fake.TokenStub = nil
	fake.tokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenFetcher) TokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.TokenStub = nil
	if fake.tokenReturnsOnCall == nil {
		fake.tokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.tokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTokenFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTokenFetcher) recordInvocation(key string, args []interface{}) {
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

var _ auth0.TokenFetcher = new(FakeTokenFetcher)
