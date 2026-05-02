package main

import "context"

var _ Interface = (*Struct)(nil) // おまじない

type Interface interface {
	Method(ctx context.Context)
}

type Struct struct{}

// ファクトリー関数
func NewStruct() Interface {
	return &Struct{}
}

// ここまで自動生成してくれる
// Method implements Interface.
func (*Struct) Method(ctx context.Context) {
	panic("unimplemented")
}
