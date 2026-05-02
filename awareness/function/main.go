package main

import (
	"context"
	"errors"
	"fmt"
)

type Struct struct{}

type Input struct{}

func (in Input) Validate() error {
	return errors.New("error")
}

type Output struct{}

// こっちよりも
// func Function(
// 	ctx context.Context,
// 	input Input,
// ) (Output, error) {
// 	if err := input.Validate(); err != nil {
// 		return Output{}, fmt.Errorf("validate input: %w", err)
// 	}

// 	return Output{}, nil
// }

// こっちが好き
func Function(
	ctx context.Context,
	input Input,
) (*Output, error) {
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("validate input: %w", err)
	}

	return &Output{}, nil
}
