package flagx

import "fmt"

type Config[T any] struct {
	Parse  func(s string) (T, error)
	String func(value T) string
}

func (c *Config[T]) Value(p *T, value T) *FlagValue[T] {
	return newFlagValue(value, p, c.Parse, c.String)
}

type FlagValue[T any] struct {
	value      *T
	parseFunc  func(s string) (T, error)
	stringFunc func(value T) string
}

func newFlagValue[T any](value T, p *T, parseFunc func(s string) (T, error), stringFunc func(value T) string) *FlagValue[T] {
	*p = value
	return &FlagValue[T]{
		value:      p,
		parseFunc:  parseFunc,
		stringFunc: stringFunc,
	}
}

func (f *FlagValue[T]) Set(s string) error {
	v, err := f.parseFunc(s)
	if err != nil {
		return err
	}
	*f.value = v
	return nil
}

func (f *FlagValue[T]) String() string {
	if f != nil && f.value != nil {
		return f.stringFunc(*f.value)
	}
	return ""
}

func (f *FlagValue[T]) Type() string {
	var v T
	return fmt.Sprintf("%T", v)
}
