// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Arturyus92/chat-server/internal/repository.MessageRepository -o message_repository_minimock.go -n MessageRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/gojuno/minimock/v3"
)

// MessageRepositoryMock implements repository.MessageRepository
type MessageRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateMessage          func(ctx context.Context, message *model.Message) (err error)
	inspectFuncCreateMessage   func(ctx context.Context, message *model.Message)
	afterCreateMessageCounter  uint64
	beforeCreateMessageCounter uint64
	CreateMessageMock          mMessageRepositoryMockCreateMessage
}

// NewMessageRepositoryMock returns a mock for repository.MessageRepository
func NewMessageRepositoryMock(t minimock.Tester) *MessageRepositoryMock {
	m := &MessageRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMessageMock = mMessageRepositoryMockCreateMessage{mock: m}
	m.CreateMessageMock.callArgs = []*MessageRepositoryMockCreateMessageParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mMessageRepositoryMockCreateMessage struct {
	mock               *MessageRepositoryMock
	defaultExpectation *MessageRepositoryMockCreateMessageExpectation
	expectations       []*MessageRepositoryMockCreateMessageExpectation

	callArgs []*MessageRepositoryMockCreateMessageParams
	mutex    sync.RWMutex
}

// MessageRepositoryMockCreateMessageExpectation specifies expectation struct of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageExpectation struct {
	mock    *MessageRepositoryMock
	params  *MessageRepositoryMockCreateMessageParams
	results *MessageRepositoryMockCreateMessageResults
	Counter uint64
}

// MessageRepositoryMockCreateMessageParams contains parameters of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageParams struct {
	ctx     context.Context
	message *model.Message
}

// MessageRepositoryMockCreateMessageResults contains results of the MessageRepository.CreateMessage
type MessageRepositoryMockCreateMessageResults struct {
	err error
}

// Expect sets up expected params for MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Expect(ctx context.Context, message *model.Message) *mMessageRepositoryMockCreateMessage {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	if mmCreateMessage.defaultExpectation == nil {
		mmCreateMessage.defaultExpectation = &MessageRepositoryMockCreateMessageExpectation{}
	}

	mmCreateMessage.defaultExpectation.params = &MessageRepositoryMockCreateMessageParams{ctx, message}
	for _, e := range mmCreateMessage.expectations {
		if minimock.Equal(e.params, mmCreateMessage.defaultExpectation.params) {
			mmCreateMessage.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateMessage.defaultExpectation.params)
		}
	}

	return mmCreateMessage
}

// Inspect accepts an inspector function that has same arguments as the MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Inspect(f func(ctx context.Context, message *model.Message)) *mMessageRepositoryMockCreateMessage {
	if mmCreateMessage.mock.inspectFuncCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("Inspect function is already set for MessageRepositoryMock.CreateMessage")
	}

	mmCreateMessage.mock.inspectFuncCreateMessage = f

	return mmCreateMessage
}

// Return sets up results that will be returned by MessageRepository.CreateMessage
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Return(err error) *MessageRepositoryMock {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	if mmCreateMessage.defaultExpectation == nil {
		mmCreateMessage.defaultExpectation = &MessageRepositoryMockCreateMessageExpectation{mock: mmCreateMessage.mock}
	}
	mmCreateMessage.defaultExpectation.results = &MessageRepositoryMockCreateMessageResults{err}
	return mmCreateMessage.mock
}

// Set uses given function f to mock the MessageRepository.CreateMessage method
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Set(f func(ctx context.Context, message *model.Message) (err error)) *MessageRepositoryMock {
	if mmCreateMessage.defaultExpectation != nil {
		mmCreateMessage.mock.t.Fatalf("Default expectation is already set for the MessageRepository.CreateMessage method")
	}

	if len(mmCreateMessage.expectations) > 0 {
		mmCreateMessage.mock.t.Fatalf("Some expectations are already set for the MessageRepository.CreateMessage method")
	}

	mmCreateMessage.mock.funcCreateMessage = f
	return mmCreateMessage.mock
}

// When sets expectation for the MessageRepository.CreateMessage which will trigger the result defined by the following
// Then helper
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) When(ctx context.Context, message *model.Message) *MessageRepositoryMockCreateMessageExpectation {
	if mmCreateMessage.mock.funcCreateMessage != nil {
		mmCreateMessage.mock.t.Fatalf("MessageRepositoryMock.CreateMessage mock is already set by Set")
	}

	expectation := &MessageRepositoryMockCreateMessageExpectation{
		mock:   mmCreateMessage.mock,
		params: &MessageRepositoryMockCreateMessageParams{ctx, message},
	}
	mmCreateMessage.expectations = append(mmCreateMessage.expectations, expectation)
	return expectation
}

// Then sets up MessageRepository.CreateMessage return parameters for the expectation previously defined by the When method
func (e *MessageRepositoryMockCreateMessageExpectation) Then(err error) *MessageRepositoryMock {
	e.results = &MessageRepositoryMockCreateMessageResults{err}
	return e.mock
}

// CreateMessage implements repository.MessageRepository
func (mmCreateMessage *MessageRepositoryMock) CreateMessage(ctx context.Context, message *model.Message) (err error) {
	mm_atomic.AddUint64(&mmCreateMessage.beforeCreateMessageCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateMessage.afterCreateMessageCounter, 1)

	if mmCreateMessage.inspectFuncCreateMessage != nil {
		mmCreateMessage.inspectFuncCreateMessage(ctx, message)
	}

	mm_params := MessageRepositoryMockCreateMessageParams{ctx, message}

	// Record call args
	mmCreateMessage.CreateMessageMock.mutex.Lock()
	mmCreateMessage.CreateMessageMock.callArgs = append(mmCreateMessage.CreateMessageMock.callArgs, &mm_params)
	mmCreateMessage.CreateMessageMock.mutex.Unlock()

	for _, e := range mmCreateMessage.CreateMessageMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreateMessage.CreateMessageMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateMessage.CreateMessageMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateMessage.CreateMessageMock.defaultExpectation.params
		mm_got := MessageRepositoryMockCreateMessageParams{ctx, message}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateMessage.t.Errorf("MessageRepositoryMock.CreateMessage got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateMessage.CreateMessageMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateMessage.t.Fatal("No results are set for the MessageRepositoryMock.CreateMessage")
		}
		return (*mm_results).err
	}
	if mmCreateMessage.funcCreateMessage != nil {
		return mmCreateMessage.funcCreateMessage(ctx, message)
	}
	mmCreateMessage.t.Fatalf("Unexpected call to MessageRepositoryMock.CreateMessage. %v %v", ctx, message)
	return
}

// CreateMessageAfterCounter returns a count of finished MessageRepositoryMock.CreateMessage invocations
func (mmCreateMessage *MessageRepositoryMock) CreateMessageAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateMessage.afterCreateMessageCounter)
}

// CreateMessageBeforeCounter returns a count of MessageRepositoryMock.CreateMessage invocations
func (mmCreateMessage *MessageRepositoryMock) CreateMessageBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateMessage.beforeCreateMessageCounter)
}

// Calls returns a list of arguments used in each call to MessageRepositoryMock.CreateMessage.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateMessage *mMessageRepositoryMockCreateMessage) Calls() []*MessageRepositoryMockCreateMessageParams {
	mmCreateMessage.mutex.RLock()

	argCopy := make([]*MessageRepositoryMockCreateMessageParams, len(mmCreateMessage.callArgs))
	copy(argCopy, mmCreateMessage.callArgs)

	mmCreateMessage.mutex.RUnlock()

	return argCopy
}

// MinimockCreateMessageDone returns true if the count of the CreateMessage invocations corresponds
// the number of defined expectations
func (m *MessageRepositoryMock) MinimockCreateMessageDone() bool {
	for _, e := range m.CreateMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateMessageCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateMessage != nil && mm_atomic.LoadUint64(&m.afterCreateMessageCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateMessageInspect logs each unmet expectation
func (m *MessageRepositoryMock) MinimockCreateMessageInspect() {
	for _, e := range m.CreateMessageMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepositoryMock.CreateMessage with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMessageMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateMessageCounter) < 1 {
		if m.CreateMessageMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageRepositoryMock.CreateMessage")
		} else {
			m.t.Errorf("Expected call to MessageRepositoryMock.CreateMessage with params: %#v", *m.CreateMessageMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateMessage != nil && mm_atomic.LoadUint64(&m.afterCreateMessageCounter) < 1 {
		m.t.Error("Expected call to MessageRepositoryMock.CreateMessage")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateMessageInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateMessageDone()
}
