// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package zap

import (
	"errors"
	"fmt"
)

var errMarshalNilLevel = errors.New("can't marshal a nil *Level to text")

// A Level is a logging priority. Higher levels are more important.
//
// Note that Level satisfies the Option interface, so any Level can be passed to
// NewJSON to override the default logging priority.
type Level int32

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
)

var (
	// The strings are the strings associated with a given Level.  This
	// enables customization of a levels string.
	DebugString = "debug"
	InfoString  = "info"
	WarnString  = "warn"
	ErrorString = "error"
	PanicString = "panic"
	FatalString = "fatal"
)

// String returns a lower-case ASCII representation of the log level.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return DebugString
	case InfoLevel:
		return InfoString
	case WarnLevel:
		return WarnString
	case ErrorLevel:
		return ErrorString
	case PanicLevel:
		return PanicString
	case FatalLevel:
		return FatalString
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

// MarshalText marshals the Level to text. Note that the text representation
// drops the -Level suffix (see example).
func (l *Level) MarshalText() ([]byte, error) {
	if l == nil {
		return nil, errMarshalNilLevel
	}
	return []byte(l.String()), nil
}

// UnmarshalText unmarshals text to a level. Like MarshalText, UnmarshalText
// expects the text representation of a Level to drop the -Level suffix (see
// example).
//
// In particular, this makes it easy to configure logging levels using YAML,
// TOML, or JSON files.
func (l *Level) UnmarshalText(text []byte) error {
	switch string(text) {
	case DebugString:
		*l = DebugLevel
	case InfoString:
		*l = InfoLevel
	case WarnString:
		*l = WarnLevel
	case ErrorString:
		*l = ErrorLevel
	case PanicString:
		*l = PanicLevel
	case FatalString:
		*l = FatalLevel
	default:
		return fmt.Errorf("unrecognized level: %v", string(text))
	}
	return nil
}
