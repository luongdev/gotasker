package internal

import "github.com/luongdev/gotasker"

type noOpLogger struct{}

func (l noOpLogger) Debug(_ string, _ ...any) {}
func (l noOpLogger) Error(_ string, _ ...any) {}
func (l noOpLogger) Info(_ string, _ ...any)  {}
func (l noOpLogger) Warn(_ string, _ ...any)  {}

var _ gotasker.Logger = (*noOpLogger)(nil)
