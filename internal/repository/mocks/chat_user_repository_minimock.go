// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Arturyus92/chat-server/internal/repository.ChatUserRepository -o chat_user_repository_minimock.go -n ChatUserRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/gojuno/minimock/v3"
)

// ChatUserRepositoryMock implements repository.ChatUserRepository
type ChatUserRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateChat          func(ctx context.Context, chatUser *model.ChatUser) (err error)
	inspectFuncCreateChat   func(ctx context.Context, chatUser *model.ChatUser)
	afterCreateChatCounter  uint64
	beforeCreateChatCounter uint64
	CreateChatMock          mChatUserRepositoryMockCreateChat
}

// NewChatUserRepositoryMock returns a mock for repository.ChatUserRepository
func NewChatUserRepositoryMock(t minimock.Tester) *ChatUserRepositoryMock {
	m := &ChatUserRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateChatMock = mChatUserRepositoryMockCreateChat{mock: m}
	m.CreateChatMock.callArgs = []*ChatUserRepositoryMockCreateChatParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mChatUserRepositoryMockCreateChat struct {
	mock               *ChatUserRepositoryMock
	defaultExpectation *ChatUserRepositoryMockCreateChatExpectation
	expectations       []*ChatUserRepositoryMockCreateChatExpectation

	callArgs []*ChatUserRepositoryMockCreateChatParams
	mutex    sync.RWMutex
}

// ChatUserRepositoryMockCreateChatExpectation specifies expectation struct of the ChatUserRepository.CreateChat
type ChatUserRepositoryMockCreateChatExpectation struct {
	mock    *ChatUserRepositoryMock
	params  *ChatUserRepositoryMockCreateChatParams
	results *ChatUserRepositoryMockCreateChatResults
	Counter uint64
}

// ChatUserRepositoryMockCreateChatParams contains parameters of the ChatUserRepository.CreateChat
type ChatUserRepositoryMockCreateChatParams struct {
	ctx      context.Context
	chatUser *model.ChatUser
}

// ChatUserRepositoryMockCreateChatResults contains results of the ChatUserRepository.CreateChat
type ChatUserRepositoryMockCreateChatResults struct {
	err error
}

// Expect sets up expected params for ChatUserRepository.CreateChat
func (mmCreateChat *mChatUserRepositoryMockCreateChat) Expect(ctx context.Context, chatUser *model.ChatUser) *mChatUserRepositoryMockCreateChat {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatUserRepositoryMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ChatUserRepositoryMockCreateChatExpectation{}
	}

	mmCreateChat.defaultExpectation.params = &ChatUserRepositoryMockCreateChatParams{ctx, chatUser}
	for _, e := range mmCreateChat.expectations {
		if minimock.Equal(e.params, mmCreateChat.defaultExpectation.params) {
			mmCreateChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateChat.defaultExpectation.params)
		}
	}

	return mmCreateChat
}

// Inspect accepts an inspector function that has same arguments as the ChatUserRepository.CreateChat
func (mmCreateChat *mChatUserRepositoryMockCreateChat) Inspect(f func(ctx context.Context, chatUser *model.ChatUser)) *mChatUserRepositoryMockCreateChat {
	if mmCreateChat.mock.inspectFuncCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("Inspect function is already set for ChatUserRepositoryMock.CreateChat")
	}

	mmCreateChat.mock.inspectFuncCreateChat = f

	return mmCreateChat
}

// Return sets up results that will be returned by ChatUserRepository.CreateChat
func (mmCreateChat *mChatUserRepositoryMockCreateChat) Return(err error) *ChatUserRepositoryMock {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatUserRepositoryMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ChatUserRepositoryMockCreateChatExpectation{mock: mmCreateChat.mock}
	}
	mmCreateChat.defaultExpectation.results = &ChatUserRepositoryMockCreateChatResults{err}
	return mmCreateChat.mock
}

// Set uses given function f to mock the ChatUserRepository.CreateChat method
func (mmCreateChat *mChatUserRepositoryMockCreateChat) Set(f func(ctx context.Context, chatUser *model.ChatUser) (err error)) *ChatUserRepositoryMock {
	if mmCreateChat.defaultExpectation != nil {
		mmCreateChat.mock.t.Fatalf("Default expectation is already set for the ChatUserRepository.CreateChat method")
	}

	if len(mmCreateChat.expectations) > 0 {
		mmCreateChat.mock.t.Fatalf("Some expectations are already set for the ChatUserRepository.CreateChat method")
	}

	mmCreateChat.mock.funcCreateChat = f
	return mmCreateChat.mock
}

// When sets expectation for the ChatUserRepository.CreateChat which will trigger the result defined by the following
// Then helper
func (mmCreateChat *mChatUserRepositoryMockCreateChat) When(ctx context.Context, chatUser *model.ChatUser) *ChatUserRepositoryMockCreateChatExpectation {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatUserRepositoryMock.CreateChat mock is already set by Set")
	}

	expectation := &ChatUserRepositoryMockCreateChatExpectation{
		mock:   mmCreateChat.mock,
		params: &ChatUserRepositoryMockCreateChatParams{ctx, chatUser},
	}
	mmCreateChat.expectations = append(mmCreateChat.expectations, expectation)
	return expectation
}

// Then sets up ChatUserRepository.CreateChat return parameters for the expectation previously defined by the When method
func (e *ChatUserRepositoryMockCreateChatExpectation) Then(err error) *ChatUserRepositoryMock {
	e.results = &ChatUserRepositoryMockCreateChatResults{err}
	return e.mock
}

// CreateChat implements repository.ChatUserRepository
func (mmCreateChat *ChatUserRepositoryMock) CreateChat(ctx context.Context, chatUser *model.ChatUser) (err error) {
	mm_atomic.AddUint64(&mmCreateChat.beforeCreateChatCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateChat.afterCreateChatCounter, 1)

	if mmCreateChat.inspectFuncCreateChat != nil {
		mmCreateChat.inspectFuncCreateChat(ctx, chatUser)
	}

	mm_params := ChatUserRepositoryMockCreateChatParams{ctx, chatUser}

	// Record call args
	mmCreateChat.CreateChatMock.mutex.Lock()
	mmCreateChat.CreateChatMock.callArgs = append(mmCreateChat.CreateChatMock.callArgs, &mm_params)
	mmCreateChat.CreateChatMock.mutex.Unlock()

	for _, e := range mmCreateChat.CreateChatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreateChat.CreateChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateChat.CreateChatMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateChat.CreateChatMock.defaultExpectation.params
		mm_got := ChatUserRepositoryMockCreateChatParams{ctx, chatUser}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateChat.t.Errorf("ChatUserRepositoryMock.CreateChat got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateChat.CreateChatMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateChat.t.Fatal("No results are set for the ChatUserRepositoryMock.CreateChat")
		}
		return (*mm_results).err
	}
	if mmCreateChat.funcCreateChat != nil {
		return mmCreateChat.funcCreateChat(ctx, chatUser)
	}
	mmCreateChat.t.Fatalf("Unexpected call to ChatUserRepositoryMock.CreateChat. %v %v", ctx, chatUser)
	return
}

// CreateChatAfterCounter returns a count of finished ChatUserRepositoryMock.CreateChat invocations
func (mmCreateChat *ChatUserRepositoryMock) CreateChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateChat.afterCreateChatCounter)
}

// CreateChatBeforeCounter returns a count of ChatUserRepositoryMock.CreateChat invocations
func (mmCreateChat *ChatUserRepositoryMock) CreateChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateChat.beforeCreateChatCounter)
}

// Calls returns a list of arguments used in each call to ChatUserRepositoryMock.CreateChat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateChat *mChatUserRepositoryMockCreateChat) Calls() []*ChatUserRepositoryMockCreateChatParams {
	mmCreateChat.mutex.RLock()

	argCopy := make([]*ChatUserRepositoryMockCreateChatParams, len(mmCreateChat.callArgs))
	copy(argCopy, mmCreateChat.callArgs)

	mmCreateChat.mutex.RUnlock()

	return argCopy
}

// MinimockCreateChatDone returns true if the count of the CreateChat invocations corresponds
// the number of defined expectations
func (m *ChatUserRepositoryMock) MinimockCreateChatDone() bool {
	for _, e := range m.CreateChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateChat != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateChatInspect logs each unmet expectation
func (m *ChatUserRepositoryMock) MinimockCreateChatInspect() {
	for _, e := range m.CreateChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatUserRepositoryMock.CreateChat with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateChatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		if m.CreateChatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatUserRepositoryMock.CreateChat")
		} else {
			m.t.Errorf("Expected call to ChatUserRepositoryMock.CreateChat with params: %#v", *m.CreateChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateChat != nil && mm_atomic.LoadUint64(&m.afterCreateChatCounter) < 1 {
		m.t.Error("Expected call to ChatUserRepositoryMock.CreateChat")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatUserRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateChatInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatUserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ChatUserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateChatDone()
}
