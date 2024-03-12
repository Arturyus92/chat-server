// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/Arturyus92/chat-server/internal/repository.LogRepository -o log_repository_minimock.go -n LogRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	model "github.com/Arturyus92/chat-server/internal/models"
	"github.com/gojuno/minimock/v3"
)

// LogRepositoryMock implements repository.LogRepository
type LogRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateLog          func(ctx context.Context, log *model.Log) (err error)
	inspectFuncCreateLog   func(ctx context.Context, log *model.Log)
	afterCreateLogCounter  uint64
	beforeCreateLogCounter uint64
	CreateLogMock          mLogRepositoryMockCreateLog
}

// NewLogRepositoryMock returns a mock for repository.LogRepository
func NewLogRepositoryMock(t minimock.Tester) *LogRepositoryMock {
	m := &LogRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateLogMock = mLogRepositoryMockCreateLog{mock: m}
	m.CreateLogMock.callArgs = []*LogRepositoryMockCreateLogParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mLogRepositoryMockCreateLog struct {
	mock               *LogRepositoryMock
	defaultExpectation *LogRepositoryMockCreateLogExpectation
	expectations       []*LogRepositoryMockCreateLogExpectation

	callArgs []*LogRepositoryMockCreateLogParams
	mutex    sync.RWMutex
}

// LogRepositoryMockCreateLogExpectation specifies expectation struct of the LogRepository.CreateLog
type LogRepositoryMockCreateLogExpectation struct {
	mock    *LogRepositoryMock
	params  *LogRepositoryMockCreateLogParams
	results *LogRepositoryMockCreateLogResults
	Counter uint64
}

// LogRepositoryMockCreateLogParams contains parameters of the LogRepository.CreateLog
type LogRepositoryMockCreateLogParams struct {
	ctx context.Context
	log *model.Log
}

// LogRepositoryMockCreateLogResults contains results of the LogRepository.CreateLog
type LogRepositoryMockCreateLogResults struct {
	err error
}

// Expect sets up expected params for LogRepository.CreateLog
func (mmCreateLog *mLogRepositoryMockCreateLog) Expect(ctx context.Context, log *model.Log) *mLogRepositoryMockCreateLog {
	if mmCreateLog.mock.funcCreateLog != nil {
		mmCreateLog.mock.t.Fatalf("LogRepositoryMock.CreateLog mock is already set by Set")
	}

	if mmCreateLog.defaultExpectation == nil {
		mmCreateLog.defaultExpectation = &LogRepositoryMockCreateLogExpectation{}
	}

	mmCreateLog.defaultExpectation.params = &LogRepositoryMockCreateLogParams{ctx, log}
	for _, e := range mmCreateLog.expectations {
		if minimock.Equal(e.params, mmCreateLog.defaultExpectation.params) {
			mmCreateLog.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateLog.defaultExpectation.params)
		}
	}

	return mmCreateLog
}

// Inspect accepts an inspector function that has same arguments as the LogRepository.CreateLog
func (mmCreateLog *mLogRepositoryMockCreateLog) Inspect(f func(ctx context.Context, log *model.Log)) *mLogRepositoryMockCreateLog {
	if mmCreateLog.mock.inspectFuncCreateLog != nil {
		mmCreateLog.mock.t.Fatalf("Inspect function is already set for LogRepositoryMock.CreateLog")
	}

	mmCreateLog.mock.inspectFuncCreateLog = f

	return mmCreateLog
}

// Return sets up results that will be returned by LogRepository.CreateLog
func (mmCreateLog *mLogRepositoryMockCreateLog) Return(err error) *LogRepositoryMock {
	if mmCreateLog.mock.funcCreateLog != nil {
		mmCreateLog.mock.t.Fatalf("LogRepositoryMock.CreateLog mock is already set by Set")
	}

	if mmCreateLog.defaultExpectation == nil {
		mmCreateLog.defaultExpectation = &LogRepositoryMockCreateLogExpectation{mock: mmCreateLog.mock}
	}
	mmCreateLog.defaultExpectation.results = &LogRepositoryMockCreateLogResults{err}
	return mmCreateLog.mock
}

// Set uses given function f to mock the LogRepository.CreateLog method
func (mmCreateLog *mLogRepositoryMockCreateLog) Set(f func(ctx context.Context, log *model.Log) (err error)) *LogRepositoryMock {
	if mmCreateLog.defaultExpectation != nil {
		mmCreateLog.mock.t.Fatalf("Default expectation is already set for the LogRepository.CreateLog method")
	}

	if len(mmCreateLog.expectations) > 0 {
		mmCreateLog.mock.t.Fatalf("Some expectations are already set for the LogRepository.CreateLog method")
	}

	mmCreateLog.mock.funcCreateLog = f
	return mmCreateLog.mock
}

// When sets expectation for the LogRepository.CreateLog which will trigger the result defined by the following
// Then helper
func (mmCreateLog *mLogRepositoryMockCreateLog) When(ctx context.Context, log *model.Log) *LogRepositoryMockCreateLogExpectation {
	if mmCreateLog.mock.funcCreateLog != nil {
		mmCreateLog.mock.t.Fatalf("LogRepositoryMock.CreateLog mock is already set by Set")
	}

	expectation := &LogRepositoryMockCreateLogExpectation{
		mock:   mmCreateLog.mock,
		params: &LogRepositoryMockCreateLogParams{ctx, log},
	}
	mmCreateLog.expectations = append(mmCreateLog.expectations, expectation)
	return expectation
}

// Then sets up LogRepository.CreateLog return parameters for the expectation previously defined by the When method
func (e *LogRepositoryMockCreateLogExpectation) Then(err error) *LogRepositoryMock {
	e.results = &LogRepositoryMockCreateLogResults{err}
	return e.mock
}

// CreateLog implements repository.LogRepository
func (mmCreateLog *LogRepositoryMock) CreateLog(ctx context.Context, log *model.Log) (err error) {
	mm_atomic.AddUint64(&mmCreateLog.beforeCreateLogCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateLog.afterCreateLogCounter, 1)

	if mmCreateLog.inspectFuncCreateLog != nil {
		mmCreateLog.inspectFuncCreateLog(ctx, log)
	}

	mm_params := LogRepositoryMockCreateLogParams{ctx, log}

	// Record call args
	mmCreateLog.CreateLogMock.mutex.Lock()
	mmCreateLog.CreateLogMock.callArgs = append(mmCreateLog.CreateLogMock.callArgs, &mm_params)
	mmCreateLog.CreateLogMock.mutex.Unlock()

	for _, e := range mmCreateLog.CreateLogMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmCreateLog.CreateLogMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateLog.CreateLogMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateLog.CreateLogMock.defaultExpectation.params
		mm_got := LogRepositoryMockCreateLogParams{ctx, log}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateLog.t.Errorf("LogRepositoryMock.CreateLog got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateLog.CreateLogMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateLog.t.Fatal("No results are set for the LogRepositoryMock.CreateLog")
		}
		return (*mm_results).err
	}
	if mmCreateLog.funcCreateLog != nil {
		return mmCreateLog.funcCreateLog(ctx, log)
	}
	mmCreateLog.t.Fatalf("Unexpected call to LogRepositoryMock.CreateLog. %v %v", ctx, log)
	return
}

// CreateLogAfterCounter returns a count of finished LogRepositoryMock.CreateLog invocations
func (mmCreateLog *LogRepositoryMock) CreateLogAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateLog.afterCreateLogCounter)
}

// CreateLogBeforeCounter returns a count of LogRepositoryMock.CreateLog invocations
func (mmCreateLog *LogRepositoryMock) CreateLogBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateLog.beforeCreateLogCounter)
}

// Calls returns a list of arguments used in each call to LogRepositoryMock.CreateLog.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateLog *mLogRepositoryMockCreateLog) Calls() []*LogRepositoryMockCreateLogParams {
	mmCreateLog.mutex.RLock()

	argCopy := make([]*LogRepositoryMockCreateLogParams, len(mmCreateLog.callArgs))
	copy(argCopy, mmCreateLog.callArgs)

	mmCreateLog.mutex.RUnlock()

	return argCopy
}

// MinimockCreateLogDone returns true if the count of the CreateLog invocations corresponds
// the number of defined expectations
func (m *LogRepositoryMock) MinimockCreateLogDone() bool {
	for _, e := range m.CreateLogMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateLogMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateLogCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateLog != nil && mm_atomic.LoadUint64(&m.afterCreateLogCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateLogInspect logs each unmet expectation
func (m *LogRepositoryMock) MinimockCreateLogInspect() {
	for _, e := range m.CreateLogMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLog with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateLogMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateLogCounter) < 1 {
		if m.CreateLogMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to LogRepositoryMock.CreateLog")
		} else {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateLog with params: %#v", *m.CreateLogMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateLog != nil && mm_atomic.LoadUint64(&m.afterCreateLogCounter) < 1 {
		m.t.Error("Expected call to LogRepositoryMock.CreateLog")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LogRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateLogInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LogRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *LogRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateLogDone()
}
