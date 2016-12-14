// This file was generated by counterfeiter
// counterfeiter -o ios/iexec/cmd_fake.go --fake-name CmdFake ios/iexec/cmd.go Cmd

package iexec

import (
	"io"
	"os"
	"sync"
	"syscall"

	"github.com/BooleanCat/igo/ios"
)

//CmdFake ...
type CmdFake struct {
	CombinedOutputStub        func() ([]byte, error)
	combinedOutputMutex       sync.RWMutex
	combinedOutputArgsForCall []struct{}
	combinedOutputReturns     struct {
		result1 []byte
		result2 error
	}
	OutputStub        func() ([]byte, error)
	outputMutex       sync.RWMutex
	outputArgsForCall []struct{}
	outputReturns     struct {
		result1 []byte
		result2 error
	}
	RunStub        func() error
	runMutex       sync.RWMutex
	runArgsForCall []struct{}
	runReturns     struct {
		result1 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	StderrPipeStub        func() (io.ReadCloser, error)
	stderrPipeMutex       sync.RWMutex
	stderrPipeArgsForCall []struct{}
	stderrPipeReturns     struct {
		result1 io.ReadCloser
		result2 error
	}
	StdinPipeStub        func() (io.WriteCloser, error)
	stdinPipeMutex       sync.RWMutex
	stdinPipeArgsForCall []struct{}
	stdinPipeReturns     struct {
		result1 io.WriteCloser
		result2 error
	}
	StdoutPipeStub        func() (io.ReadCloser, error)
	stdoutPipeMutex       sync.RWMutex
	stdoutPipeArgsForCall []struct{}
	stdoutPipeReturns     struct {
		result1 io.ReadCloser
		result2 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
	GetPathStub        func() string
	getPathMutex       sync.RWMutex
	getPathArgsForCall []struct{}
	getPathReturns     struct {
		result1 string
	}
	SetPathStub        func(string)
	setPathMutex       sync.RWMutex
	setPathArgsForCall []struct {
		arg1 string
	}
	GetArgsStub        func() []string
	getArgsMutex       sync.RWMutex
	getArgsArgsForCall []struct{}
	getArgsReturns     struct {
		result1 []string
	}
	SetArgsStub        func([]string)
	setArgsMutex       sync.RWMutex
	setArgsArgsForCall []struct {
		arg1 []string
	}
	GetEnvStub        func() []string
	getEnvMutex       sync.RWMutex
	getEnvArgsForCall []struct{}
	getEnvReturns     struct {
		result1 []string
	}
	SetEnvStub        func([]string)
	setEnvMutex       sync.RWMutex
	setEnvArgsForCall []struct {
		arg1 []string
	}
	GetDirStub        func() string
	getDirMutex       sync.RWMutex
	getDirArgsForCall []struct{}
	getDirReturns     struct {
		result1 string
	}
	SetDirStub        func(string)
	setDirMutex       sync.RWMutex
	setDirArgsForCall []struct {
		arg1 string
	}
	GetStdinStub        func() io.Reader
	getStdinMutex       sync.RWMutex
	getStdinArgsForCall []struct{}
	getStdinReturns     struct {
		result1 io.Reader
	}
	SetStdinStub        func(io.Reader)
	setStdinMutex       sync.RWMutex
	setStdinArgsForCall []struct {
		arg1 io.Reader
	}
	GetStdoutStub        func() io.Writer
	getStdoutMutex       sync.RWMutex
	getStdoutArgsForCall []struct{}
	getStdoutReturns     struct {
		result1 io.Writer
	}
	SetStdoutStub        func(io.Writer)
	setStdoutMutex       sync.RWMutex
	setStdoutArgsForCall []struct {
		arg1 io.Writer
	}
	GetStderrStub        func() io.Writer
	getStderrMutex       sync.RWMutex
	getStderrArgsForCall []struct{}
	getStderrReturns     struct {
		result1 io.Writer
	}
	SetStderrStub        func(io.Writer)
	setStderrMutex       sync.RWMutex
	setStderrArgsForCall []struct {
		arg1 io.Writer
	}
	GetExtraFilesStub        func() []*os.File
	getExtraFilesMutex       sync.RWMutex
	getExtraFilesArgsForCall []struct{}
	getExtraFilesReturns     struct {
		result1 []*os.File
	}
	SetExtraFilesStub        func([]*os.File)
	setExtraFilesMutex       sync.RWMutex
	setExtraFilesArgsForCall []struct {
		arg1 []*os.File
	}
	GetSysProcAttrStub        func() *syscall.SysProcAttr
	getSysProcAttrMutex       sync.RWMutex
	getSysProcAttrArgsForCall []struct{}
	getSysProcAttrReturns     struct {
		result1 *syscall.SysProcAttr
	}
	SetSysProcAttrStub        func(*syscall.SysProcAttr)
	setSysProcAttrMutex       sync.RWMutex
	setSysProcAttrArgsForCall []struct {
		arg1 *syscall.SysProcAttr
	}
	GetProcessStub        func() ios.Process
	getProcessMutex       sync.RWMutex
	getProcessArgsForCall []struct{}
	getProcessReturns     struct {
		result1 ios.Process
	}
	SetProcessStub        func(*os.Process)
	setProcessMutex       sync.RWMutex
	setProcessArgsForCall []struct {
		arg1 *os.Process
	}
	GetProcessStateStub        func() *os.ProcessState
	getProcessStateMutex       sync.RWMutex
	getProcessStateArgsForCall []struct{}
	getProcessStateReturns     struct {
		result1 *os.ProcessState
	}
	SetProcessStateStub        func(*os.ProcessState)
	setProcessStateMutex       sync.RWMutex
	setProcessStateArgsForCall []struct {
		arg1 *os.ProcessState
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

//CombinedOutput ...
func (fake *CmdFake) CombinedOutput() ([]byte, error) {
	fake.combinedOutputMutex.Lock()
	fake.combinedOutputArgsForCall = append(fake.combinedOutputArgsForCall, struct{}{})
	fake.recordInvocation("CombinedOutput", []interface{}{})
	fake.combinedOutputMutex.Unlock()
	if fake.CombinedOutputStub != nil {
		return fake.CombinedOutputStub()
	}
	return fake.combinedOutputReturns.result1, fake.combinedOutputReturns.result2
}

//CombinedOutputCallCount ...
func (fake *CmdFake) CombinedOutputCallCount() int {
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	return len(fake.combinedOutputArgsForCall)
}

//CombinedOutputReturns ...
func (fake *CmdFake) CombinedOutputReturns(result1 []byte, result2 error) {
	fake.CombinedOutputStub = nil
	fake.combinedOutputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

//Output ...
func (fake *CmdFake) Output() ([]byte, error) {
	fake.outputMutex.Lock()
	fake.outputArgsForCall = append(fake.outputArgsForCall, struct{}{})
	fake.recordInvocation("Output", []interface{}{})
	fake.outputMutex.Unlock()
	if fake.OutputStub != nil {
		return fake.OutputStub()
	}
	return fake.outputReturns.result1, fake.outputReturns.result2
}

//OutputCallCount ...
func (fake *CmdFake) OutputCallCount() int {
	fake.outputMutex.RLock()
	defer fake.outputMutex.RUnlock()
	return len(fake.outputArgsForCall)
}

//OutputReturns ...
func (fake *CmdFake) OutputReturns(result1 []byte, result2 error) {
	fake.OutputStub = nil
	fake.outputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

//Run ...
func (fake *CmdFake) Run() error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct{}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub()
	}
	return fake.runReturns.result1
}

//RunCallCount ...
func (fake *CmdFake) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

//RunReturns ...
func (fake *CmdFake) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

//Start ...
func (fake *CmdFake) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	return fake.startReturns.result1
}

//StartCallCount ...
func (fake *CmdFake) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

//StartReturns ...
func (fake *CmdFake) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

//StderrPipe ...
func (fake *CmdFake) StderrPipe() (io.ReadCloser, error) {
	fake.stderrPipeMutex.Lock()
	fake.stderrPipeArgsForCall = append(fake.stderrPipeArgsForCall, struct{}{})
	fake.recordInvocation("StderrPipe", []interface{}{})
	fake.stderrPipeMutex.Unlock()
	if fake.StderrPipeStub != nil {
		return fake.StderrPipeStub()
	}
	return fake.stderrPipeReturns.result1, fake.stderrPipeReturns.result2
}

//StderrPipeCallCount ...
func (fake *CmdFake) StderrPipeCallCount() int {
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	return len(fake.stderrPipeArgsForCall)
}

//StderrPipeReturns ...
func (fake *CmdFake) StderrPipeReturns(result1 io.ReadCloser, result2 error) {
	fake.StderrPipeStub = nil
	fake.stderrPipeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

//StdinPipe ...
func (fake *CmdFake) StdinPipe() (io.WriteCloser, error) {
	fake.stdinPipeMutex.Lock()
	fake.stdinPipeArgsForCall = append(fake.stdinPipeArgsForCall, struct{}{})
	fake.recordInvocation("StdinPipe", []interface{}{})
	fake.stdinPipeMutex.Unlock()
	if fake.StdinPipeStub != nil {
		return fake.StdinPipeStub()
	}
	return fake.stdinPipeReturns.result1, fake.stdinPipeReturns.result2
}

//StdinPipeCallCount ...
func (fake *CmdFake) StdinPipeCallCount() int {
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	return len(fake.stdinPipeArgsForCall)
}

//StdinPipeReturns ...
func (fake *CmdFake) StdinPipeReturns(result1 io.WriteCloser, result2 error) {
	fake.StdinPipeStub = nil
	fake.stdinPipeReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

//StdoutPipe ...
func (fake *CmdFake) StdoutPipe() (io.ReadCloser, error) {
	fake.stdoutPipeMutex.Lock()
	fake.stdoutPipeArgsForCall = append(fake.stdoutPipeArgsForCall, struct{}{})
	fake.recordInvocation("StdoutPipe", []interface{}{})
	fake.stdoutPipeMutex.Unlock()
	if fake.StdoutPipeStub != nil {
		return fake.StdoutPipeStub()
	}
	return fake.stdoutPipeReturns.result1, fake.stdoutPipeReturns.result2
}

//StdoutPipeCallCount ...
func (fake *CmdFake) StdoutPipeCallCount() int {
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	return len(fake.stdoutPipeArgsForCall)
}

//StdoutPipeReturns ...
func (fake *CmdFake) StdoutPipeReturns(result1 io.ReadCloser, result2 error) {
	fake.StdoutPipeStub = nil
	fake.stdoutPipeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

//Wait ...
func (fake *CmdFake) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	}
	return fake.waitReturns.result1
}

//WaitCallCount ...
func (fake *CmdFake) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

//WaitReturns ...
func (fake *CmdFake) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

//GetPath ...
func (fake *CmdFake) GetPath() string {
	fake.getPathMutex.Lock()
	fake.getPathArgsForCall = append(fake.getPathArgsForCall, struct{}{})
	fake.recordInvocation("GetPath", []interface{}{})
	fake.getPathMutex.Unlock()
	if fake.GetPathStub != nil {
		return fake.GetPathStub()
	}
	return fake.getPathReturns.result1
}

//GetPathCallCount ...
func (fake *CmdFake) GetPathCallCount() int {
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	return len(fake.getPathArgsForCall)
}

//GetPathReturns ...
func (fake *CmdFake) GetPathReturns(result1 string) {
	fake.GetPathStub = nil
	fake.getPathReturns = struct {
		result1 string
	}{result1}
}

//SetPath ...
func (fake *CmdFake) SetPath(arg1 string) {
	fake.setPathMutex.Lock()
	fake.setPathArgsForCall = append(fake.setPathArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetPath", []interface{}{arg1})
	fake.setPathMutex.Unlock()
	if fake.SetPathStub != nil {
		fake.SetPathStub(arg1)
	}
}

//SetPathCallCount ...
func (fake *CmdFake) SetPathCallCount() int {
	fake.setPathMutex.RLock()
	defer fake.setPathMutex.RUnlock()
	return len(fake.setPathArgsForCall)
}

//SetPathArgsForCall ...
func (fake *CmdFake) SetPathArgsForCall(i int) string {
	fake.setPathMutex.RLock()
	defer fake.setPathMutex.RUnlock()
	return fake.setPathArgsForCall[i].arg1
}

//GetArgs ...
func (fake *CmdFake) GetArgs() []string {
	fake.getArgsMutex.Lock()
	fake.getArgsArgsForCall = append(fake.getArgsArgsForCall, struct{}{})
	fake.recordInvocation("GetArgs", []interface{}{})
	fake.getArgsMutex.Unlock()
	if fake.GetArgsStub != nil {
		return fake.GetArgsStub()
	}
	return fake.getArgsReturns.result1
}

//GetArgsCallCount ...
func (fake *CmdFake) GetArgsCallCount() int {
	fake.getArgsMutex.RLock()
	defer fake.getArgsMutex.RUnlock()
	return len(fake.getArgsArgsForCall)
}

//GetArgsReturns ...
func (fake *CmdFake) GetArgsReturns(result1 []string) {
	fake.GetArgsStub = nil
	fake.getArgsReturns = struct {
		result1 []string
	}{result1}
}

//SetArgs ...
func (fake *CmdFake) SetArgs(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setArgsMutex.Lock()
	fake.setArgsArgsForCall = append(fake.setArgsArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("SetArgs", []interface{}{arg1Copy})
	fake.setArgsMutex.Unlock()
	if fake.SetArgsStub != nil {
		fake.SetArgsStub(arg1)
	}
}

//SetArgsCallCount ...
func (fake *CmdFake) SetArgsCallCount() int {
	fake.setArgsMutex.RLock()
	defer fake.setArgsMutex.RUnlock()
	return len(fake.setArgsArgsForCall)
}

//SetArgsArgsForCall ...
func (fake *CmdFake) SetArgsArgsForCall(i int) []string {
	fake.setArgsMutex.RLock()
	defer fake.setArgsMutex.RUnlock()
	return fake.setArgsArgsForCall[i].arg1
}

//GetEnv ...
func (fake *CmdFake) GetEnv() []string {
	fake.getEnvMutex.Lock()
	fake.getEnvArgsForCall = append(fake.getEnvArgsForCall, struct{}{})
	fake.recordInvocation("GetEnv", []interface{}{})
	fake.getEnvMutex.Unlock()
	if fake.GetEnvStub != nil {
		return fake.GetEnvStub()
	}
	return fake.getEnvReturns.result1
}

//GetEnvCallCount ...
func (fake *CmdFake) GetEnvCallCount() int {
	fake.getEnvMutex.RLock()
	defer fake.getEnvMutex.RUnlock()
	return len(fake.getEnvArgsForCall)
}

//GetEnvReturns ...
func (fake *CmdFake) GetEnvReturns(result1 []string) {
	fake.GetEnvStub = nil
	fake.getEnvReturns = struct {
		result1 []string
	}{result1}
}

//SetEnv ...
func (fake *CmdFake) SetEnv(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setEnvMutex.Lock()
	fake.setEnvArgsForCall = append(fake.setEnvArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("SetEnv", []interface{}{arg1Copy})
	fake.setEnvMutex.Unlock()
	if fake.SetEnvStub != nil {
		fake.SetEnvStub(arg1)
	}
}

//SetEnvCallCount ...
func (fake *CmdFake) SetEnvCallCount() int {
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	return len(fake.setEnvArgsForCall)
}

//SetEnvArgsForCall ...
func (fake *CmdFake) SetEnvArgsForCall(i int) []string {
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	return fake.setEnvArgsForCall[i].arg1
}

//GetDir ...
func (fake *CmdFake) GetDir() string {
	fake.getDirMutex.Lock()
	fake.getDirArgsForCall = append(fake.getDirArgsForCall, struct{}{})
	fake.recordInvocation("GetDir", []interface{}{})
	fake.getDirMutex.Unlock()
	if fake.GetDirStub != nil {
		return fake.GetDirStub()
	}
	return fake.getDirReturns.result1
}

//GetDirCallCount ...
func (fake *CmdFake) GetDirCallCount() int {
	fake.getDirMutex.RLock()
	defer fake.getDirMutex.RUnlock()
	return len(fake.getDirArgsForCall)
}

//GetDirReturns ...
func (fake *CmdFake) GetDirReturns(result1 string) {
	fake.GetDirStub = nil
	fake.getDirReturns = struct {
		result1 string
	}{result1}
}

//SetDir ...
func (fake *CmdFake) SetDir(arg1 string) {
	fake.setDirMutex.Lock()
	fake.setDirArgsForCall = append(fake.setDirArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetDir", []interface{}{arg1})
	fake.setDirMutex.Unlock()
	if fake.SetDirStub != nil {
		fake.SetDirStub(arg1)
	}
}

//SetDirCallCount ...
func (fake *CmdFake) SetDirCallCount() int {
	fake.setDirMutex.RLock()
	defer fake.setDirMutex.RUnlock()
	return len(fake.setDirArgsForCall)
}

//SetDirArgsForCall ...
func (fake *CmdFake) SetDirArgsForCall(i int) string {
	fake.setDirMutex.RLock()
	defer fake.setDirMutex.RUnlock()
	return fake.setDirArgsForCall[i].arg1
}

//GetStdin ...
func (fake *CmdFake) GetStdin() io.Reader {
	fake.getStdinMutex.Lock()
	fake.getStdinArgsForCall = append(fake.getStdinArgsForCall, struct{}{})
	fake.recordInvocation("GetStdin", []interface{}{})
	fake.getStdinMutex.Unlock()
	if fake.GetStdinStub != nil {
		return fake.GetStdinStub()
	}
	return fake.getStdinReturns.result1
}

//GetStdinCallCount ...
func (fake *CmdFake) GetStdinCallCount() int {
	fake.getStdinMutex.RLock()
	defer fake.getStdinMutex.RUnlock()
	return len(fake.getStdinArgsForCall)
}

//GetStdinReturns ...
func (fake *CmdFake) GetStdinReturns(result1 io.Reader) {
	fake.GetStdinStub = nil
	fake.getStdinReturns = struct {
		result1 io.Reader
	}{result1}
}

//SetStdin ...
func (fake *CmdFake) SetStdin(arg1 io.Reader) {
	fake.setStdinMutex.Lock()
	fake.setStdinArgsForCall = append(fake.setStdinArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	fake.recordInvocation("SetStdin", []interface{}{arg1})
	fake.setStdinMutex.Unlock()
	if fake.SetStdinStub != nil {
		fake.SetStdinStub(arg1)
	}
}

//SetStdinCallCount ...
func (fake *CmdFake) SetStdinCallCount() int {
	fake.setStdinMutex.RLock()
	defer fake.setStdinMutex.RUnlock()
	return len(fake.setStdinArgsForCall)
}

//SetStdinArgsForCall ...
func (fake *CmdFake) SetStdinArgsForCall(i int) io.Reader {
	fake.setStdinMutex.RLock()
	defer fake.setStdinMutex.RUnlock()
	return fake.setStdinArgsForCall[i].arg1
}

//GetStdout ...
func (fake *CmdFake) GetStdout() io.Writer {
	fake.getStdoutMutex.Lock()
	fake.getStdoutArgsForCall = append(fake.getStdoutArgsForCall, struct{}{})
	fake.recordInvocation("GetStdout", []interface{}{})
	fake.getStdoutMutex.Unlock()
	if fake.GetStdoutStub != nil {
		return fake.GetStdoutStub()
	}
	return fake.getStdoutReturns.result1
}

//GetStdoutCallCount ...
func (fake *CmdFake) GetStdoutCallCount() int {
	fake.getStdoutMutex.RLock()
	defer fake.getStdoutMutex.RUnlock()
	return len(fake.getStdoutArgsForCall)
}

//GetStdoutReturns ...
func (fake *CmdFake) GetStdoutReturns(result1 io.Writer) {
	fake.GetStdoutStub = nil
	fake.getStdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

//SetStdout ...
func (fake *CmdFake) SetStdout(arg1 io.Writer) {
	fake.setStdoutMutex.Lock()
	fake.setStdoutArgsForCall = append(fake.setStdoutArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.recordInvocation("SetStdout", []interface{}{arg1})
	fake.setStdoutMutex.Unlock()
	if fake.SetStdoutStub != nil {
		fake.SetStdoutStub(arg1)
	}
}

//SetStdoutCallCount ...
func (fake *CmdFake) SetStdoutCallCount() int {
	fake.setStdoutMutex.RLock()
	defer fake.setStdoutMutex.RUnlock()
	return len(fake.setStdoutArgsForCall)
}

//SetStdoutArgsForCall ...
func (fake *CmdFake) SetStdoutArgsForCall(i int) io.Writer {
	fake.setStdoutMutex.RLock()
	defer fake.setStdoutMutex.RUnlock()
	return fake.setStdoutArgsForCall[i].arg1
}

//GetStderr ...
func (fake *CmdFake) GetStderr() io.Writer {
	fake.getStderrMutex.Lock()
	fake.getStderrArgsForCall = append(fake.getStderrArgsForCall, struct{}{})
	fake.recordInvocation("GetStderr", []interface{}{})
	fake.getStderrMutex.Unlock()
	if fake.GetStderrStub != nil {
		return fake.GetStderrStub()
	}
	return fake.getStderrReturns.result1
}

//GetStderrCallCount ...
func (fake *CmdFake) GetStderrCallCount() int {
	fake.getStderrMutex.RLock()
	defer fake.getStderrMutex.RUnlock()
	return len(fake.getStderrArgsForCall)
}

//GetStderrReturns ...
func (fake *CmdFake) GetStderrReturns(result1 io.Writer) {
	fake.GetStderrStub = nil
	fake.getStderrReturns = struct {
		result1 io.Writer
	}{result1}
}

//SetStderr ...
func (fake *CmdFake) SetStderr(arg1 io.Writer) {
	fake.setStderrMutex.Lock()
	fake.setStderrArgsForCall = append(fake.setStderrArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	fake.recordInvocation("SetStderr", []interface{}{arg1})
	fake.setStderrMutex.Unlock()
	if fake.SetStderrStub != nil {
		fake.SetStderrStub(arg1)
	}
}

//SetStderrCallCount ...
func (fake *CmdFake) SetStderrCallCount() int {
	fake.setStderrMutex.RLock()
	defer fake.setStderrMutex.RUnlock()
	return len(fake.setStderrArgsForCall)
}

//SetStderrArgsForCall ...
func (fake *CmdFake) SetStderrArgsForCall(i int) io.Writer {
	fake.setStderrMutex.RLock()
	defer fake.setStderrMutex.RUnlock()
	return fake.setStderrArgsForCall[i].arg1
}

//GetExtraFiles ...
func (fake *CmdFake) GetExtraFiles() []*os.File {
	fake.getExtraFilesMutex.Lock()
	fake.getExtraFilesArgsForCall = append(fake.getExtraFilesArgsForCall, struct{}{})
	fake.recordInvocation("GetExtraFiles", []interface{}{})
	fake.getExtraFilesMutex.Unlock()
	if fake.GetExtraFilesStub != nil {
		return fake.GetExtraFilesStub()
	}
	return fake.getExtraFilesReturns.result1
}

//GetExtraFilesCallCount ...
func (fake *CmdFake) GetExtraFilesCallCount() int {
	fake.getExtraFilesMutex.RLock()
	defer fake.getExtraFilesMutex.RUnlock()
	return len(fake.getExtraFilesArgsForCall)
}

//GetExtraFilesReturns ...
func (fake *CmdFake) GetExtraFilesReturns(result1 []*os.File) {
	fake.GetExtraFilesStub = nil
	fake.getExtraFilesReturns = struct {
		result1 []*os.File
	}{result1}
}

//SetExtraFiles ...
func (fake *CmdFake) SetExtraFiles(arg1 []*os.File) {
	var arg1Copy []*os.File
	if arg1 != nil {
		arg1Copy = make([]*os.File, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setExtraFilesMutex.Lock()
	fake.setExtraFilesArgsForCall = append(fake.setExtraFilesArgsForCall, struct {
		arg1 []*os.File
	}{arg1Copy})
	fake.recordInvocation("SetExtraFiles", []interface{}{arg1Copy})
	fake.setExtraFilesMutex.Unlock()
	if fake.SetExtraFilesStub != nil {
		fake.SetExtraFilesStub(arg1)
	}
}

//SetExtraFilesCallCount ...
func (fake *CmdFake) SetExtraFilesCallCount() int {
	fake.setExtraFilesMutex.RLock()
	defer fake.setExtraFilesMutex.RUnlock()
	return len(fake.setExtraFilesArgsForCall)
}

//SetExtraFilesArgsForCall ...
func (fake *CmdFake) SetExtraFilesArgsForCall(i int) []*os.File {
	fake.setExtraFilesMutex.RLock()
	defer fake.setExtraFilesMutex.RUnlock()
	return fake.setExtraFilesArgsForCall[i].arg1
}

//GetSysProcAttr ...
func (fake *CmdFake) GetSysProcAttr() *syscall.SysProcAttr {
	fake.getSysProcAttrMutex.Lock()
	fake.getSysProcAttrArgsForCall = append(fake.getSysProcAttrArgsForCall, struct{}{})
	fake.recordInvocation("GetSysProcAttr", []interface{}{})
	fake.getSysProcAttrMutex.Unlock()
	if fake.GetSysProcAttrStub != nil {
		return fake.GetSysProcAttrStub()
	}
	return fake.getSysProcAttrReturns.result1
}

//GetSysProcAttrCallCount ...
func (fake *CmdFake) GetSysProcAttrCallCount() int {
	fake.getSysProcAttrMutex.RLock()
	defer fake.getSysProcAttrMutex.RUnlock()
	return len(fake.getSysProcAttrArgsForCall)
}

//GetSysProcAttrReturns ...
func (fake *CmdFake) GetSysProcAttrReturns(result1 *syscall.SysProcAttr) {
	fake.GetSysProcAttrStub = nil
	fake.getSysProcAttrReturns = struct {
		result1 *syscall.SysProcAttr
	}{result1}
}

//SetSysProcAttr ...
func (fake *CmdFake) SetSysProcAttr(arg1 *syscall.SysProcAttr) {
	fake.setSysProcAttrMutex.Lock()
	fake.setSysProcAttrArgsForCall = append(fake.setSysProcAttrArgsForCall, struct {
		arg1 *syscall.SysProcAttr
	}{arg1})
	fake.recordInvocation("SetSysProcAttr", []interface{}{arg1})
	fake.setSysProcAttrMutex.Unlock()
	if fake.SetSysProcAttrStub != nil {
		fake.SetSysProcAttrStub(arg1)
	}
}

//SetSysProcAttrCallCount ...
func (fake *CmdFake) SetSysProcAttrCallCount() int {
	fake.setSysProcAttrMutex.RLock()
	defer fake.setSysProcAttrMutex.RUnlock()
	return len(fake.setSysProcAttrArgsForCall)
}

//SetSysProcAttrArgsForCall ...
func (fake *CmdFake) SetSysProcAttrArgsForCall(i int) *syscall.SysProcAttr {
	fake.setSysProcAttrMutex.RLock()
	defer fake.setSysProcAttrMutex.RUnlock()
	return fake.setSysProcAttrArgsForCall[i].arg1
}

//GetProcess ...
func (fake *CmdFake) GetProcess() ios.Process {
	fake.getProcessMutex.Lock()
	fake.getProcessArgsForCall = append(fake.getProcessArgsForCall, struct{}{})
	fake.recordInvocation("GetProcess", []interface{}{})
	fake.getProcessMutex.Unlock()
	if fake.GetProcessStub != nil {
		return fake.GetProcessStub()
	}
	return fake.getProcessReturns.result1
}

//GetProcessCallCount ...
func (fake *CmdFake) GetProcessCallCount() int {
	fake.getProcessMutex.RLock()
	defer fake.getProcessMutex.RUnlock()
	return len(fake.getProcessArgsForCall)
}

//GetProcessReturns ...
func (fake *CmdFake) GetProcessReturns(result1 ios.Process) {
	fake.GetProcessStub = nil
	fake.getProcessReturns = struct {
		result1 ios.Process
	}{result1}
}

//SetProcess ...
func (fake *CmdFake) SetProcess(arg1 *os.Process) {
	fake.setProcessMutex.Lock()
	fake.setProcessArgsForCall = append(fake.setProcessArgsForCall, struct {
		arg1 *os.Process
	}{arg1})
	fake.recordInvocation("SetProcess", []interface{}{arg1})
	fake.setProcessMutex.Unlock()
	if fake.SetProcessStub != nil {
		fake.SetProcessStub(arg1)
	}
}

//SetProcessCallCount ...
func (fake *CmdFake) SetProcessCallCount() int {
	fake.setProcessMutex.RLock()
	defer fake.setProcessMutex.RUnlock()
	return len(fake.setProcessArgsForCall)
}

//SetProcessArgsForCall ...
func (fake *CmdFake) SetProcessArgsForCall(i int) *os.Process {
	fake.setProcessMutex.RLock()
	defer fake.setProcessMutex.RUnlock()
	return fake.setProcessArgsForCall[i].arg1
}

//GetProcessState ...
func (fake *CmdFake) GetProcessState() *os.ProcessState {
	fake.getProcessStateMutex.Lock()
	fake.getProcessStateArgsForCall = append(fake.getProcessStateArgsForCall, struct{}{})
	fake.recordInvocation("GetProcessState", []interface{}{})
	fake.getProcessStateMutex.Unlock()
	if fake.GetProcessStateStub != nil {
		return fake.GetProcessStateStub()
	}
	return fake.getProcessStateReturns.result1
}

//GetProcessStateCallCount ...
func (fake *CmdFake) GetProcessStateCallCount() int {
	fake.getProcessStateMutex.RLock()
	defer fake.getProcessStateMutex.RUnlock()
	return len(fake.getProcessStateArgsForCall)
}

//GetProcessStateReturns ...
func (fake *CmdFake) GetProcessStateReturns(result1 *os.ProcessState) {
	fake.GetProcessStateStub = nil
	fake.getProcessStateReturns = struct {
		result1 *os.ProcessState
	}{result1}
}

//SetProcessState ...
func (fake *CmdFake) SetProcessState(arg1 *os.ProcessState) {
	fake.setProcessStateMutex.Lock()
	fake.setProcessStateArgsForCall = append(fake.setProcessStateArgsForCall, struct {
		arg1 *os.ProcessState
	}{arg1})
	fake.recordInvocation("SetProcessState", []interface{}{arg1})
	fake.setProcessStateMutex.Unlock()
	if fake.SetProcessStateStub != nil {
		fake.SetProcessStateStub(arg1)
	}
}

//SetProcessStateCallCount ...
func (fake *CmdFake) SetProcessStateCallCount() int {
	fake.setProcessStateMutex.RLock()
	defer fake.setProcessStateMutex.RUnlock()
	return len(fake.setProcessStateArgsForCall)
}

//SetProcessStateArgsForCall ...
func (fake *CmdFake) SetProcessStateArgsForCall(i int) *os.ProcessState {
	fake.setProcessStateMutex.RLock()
	defer fake.setProcessStateMutex.RUnlock()
	return fake.setProcessStateArgsForCall[i].arg1
}

//Invocations ...
func (fake *CmdFake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	fake.outputMutex.RLock()
	defer fake.outputMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	fake.getPathMutex.RLock()
	defer fake.getPathMutex.RUnlock()
	fake.setPathMutex.RLock()
	defer fake.setPathMutex.RUnlock()
	fake.getArgsMutex.RLock()
	defer fake.getArgsMutex.RUnlock()
	fake.setArgsMutex.RLock()
	defer fake.setArgsMutex.RUnlock()
	fake.getEnvMutex.RLock()
	defer fake.getEnvMutex.RUnlock()
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	fake.getDirMutex.RLock()
	defer fake.getDirMutex.RUnlock()
	fake.setDirMutex.RLock()
	defer fake.setDirMutex.RUnlock()
	fake.getStdinMutex.RLock()
	defer fake.getStdinMutex.RUnlock()
	fake.setStdinMutex.RLock()
	defer fake.setStdinMutex.RUnlock()
	fake.getStdoutMutex.RLock()
	defer fake.getStdoutMutex.RUnlock()
	fake.setStdoutMutex.RLock()
	defer fake.setStdoutMutex.RUnlock()
	fake.getStderrMutex.RLock()
	defer fake.getStderrMutex.RUnlock()
	fake.setStderrMutex.RLock()
	defer fake.setStderrMutex.RUnlock()
	fake.getExtraFilesMutex.RLock()
	defer fake.getExtraFilesMutex.RUnlock()
	fake.setExtraFilesMutex.RLock()
	defer fake.setExtraFilesMutex.RUnlock()
	fake.getSysProcAttrMutex.RLock()
	defer fake.getSysProcAttrMutex.RUnlock()
	fake.setSysProcAttrMutex.RLock()
	defer fake.setSysProcAttrMutex.RUnlock()
	fake.getProcessMutex.RLock()
	defer fake.getProcessMutex.RUnlock()
	fake.setProcessMutex.RLock()
	defer fake.setProcessMutex.RUnlock()
	fake.getProcessStateMutex.RLock()
	defer fake.getProcessStateMutex.RUnlock()
	fake.setProcessStateMutex.RLock()
	defer fake.setProcessStateMutex.RUnlock()
	return fake.invocations
}

func (fake *CmdFake) recordInvocation(key string, args []interface{}) {
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

var _ Cmd = new(CmdFake)
