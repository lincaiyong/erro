package erro

import (
	"errors"
	"fmt"
	"github.com/lincaiyong/log"
	"runtime"
	"strings"
)

func Recover(fn func(error)) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			fn(err)
		} else {
			panic(r)
		}
	}
}

func traceMsg(msg string) string {
	const maxStackDepth = 32
	pcs := make([]uintptr, maxStackDepth)
	n := runtime.Callers(3, pcs) // 2 = skip runtime.Callers + this + Assert

	frames := runtime.CallersFrames(pcs[:n])
	var sb strings.Builder
	sb.WriteString(msg)
	sb.WriteString("\nStack trace:\n")
	for {
		frame, more := frames.Next()
		if !strings.HasPrefix(frame.Function, "runtime.") {
			sb.WriteString(fmt.Sprintf("  %s ( %s:%d )\n", frame.Function, frame.File, frame.Line))
		}
		if !more {
			break
		}
	}
	return sb.String()
}

func Assert(b bool, msg string, args ...any) {
	if !b {
		if len(args) > 0 {
			msg = fmt.Sprintf(msg, args...)
		}
		msg = traceMsg(msg)
		panic(errors.New(msg))
	}
}

func checkInfo(err error, msg string, args ...any) {
	if err != nil {
		if len(args) > 0 {
			msg = fmt.Sprintf(msg, args...)
		}
		msg = traceMsg(fmt.Sprintf("%s: %v", msg, err))
		panic(errors.New(msg))
	} else {
		log.InfoLog(msg, args...)
	}
}

func checkDebug(err error, msg string, args ...any) {
	if err != nil {
		if len(args) > 0 {
			msg = fmt.Sprintf(msg, args...)
		}
		msg = traceMsg(fmt.Sprintf("%s: %v", msg, err))
		panic(errors.New(msg))
	} else {
		log.DebugLog(msg, args...)
	}
}

type C0 struct {
	err error
}

func E0(err error) *C0 {
	return &C0{err: err}
}

func (c *C0) Info(msg string, args ...any) {
	checkInfo(c.err, msg, args...)
}

func (c *C0) Debug(msg string, args ...any) {
	checkDebug(c.err, msg, args...)
}

type C1[T any] struct {
	v   T
	err error
}

func E1[T any](v T, err error) *C1[T] {
	return &C1[T]{v: v, err: err}
}

func (c *C1[T]) Info(msg string, args ...any) T {
	checkInfo(c.err, msg, args...)
	return c.v
}

func (c *C1[T]) Debug(msg string, args ...any) T {
	checkDebug(c.err, msg, args...)
	return c.v
}

type C2[T1, T2 any] struct {
	v1  T1
	v2  T2
	err error
}

func E2[T1, T2 any](v1 T1, v2 T2, err error) *C2[T1, T2] {
	return &C2[T1, T2]{v1: v1, v2: v2, err: err}
}

func (c *C2[T1, T2]) Info(msg string, args ...any) (T1, T2) {
	checkInfo(c.err, msg, args...)
	return c.v1, c.v2
}

func (c *C2[T1, T2]) Debug(msg string, args ...any) (T1, T2) {
	checkDebug(c.err, msg, args...)
	return c.v1, c.v2
}

type C3[T1, T2, T3 any] struct {
	v1  T1
	v2  T2
	v3  T3
	err error
}

func E3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) *C3[T1, T2, T3] {
	return &C3[T1, T2, T3]{v1: v1, v2: v2, v3: v3, err: err}
}

func (c *C3[T1, T2, T3]) Info(msg string, args ...any) (T1, T2, T3) {
	checkInfo(c.err, msg, args...)
	return c.v1, c.v2, c.v3
}

func (c *C3[T1, T2, T3]) Debug(msg string, args ...any) (T1, T2, T3) {
	checkDebug(c.err, msg, args...)
	return c.v1, c.v2, c.v3
}
