package main

import "errors"

type Result struct {
	Value int
	Err   error
}

func Process(value int) Result {
	switch {
	case value < 0:
		return Result{
			Value: value,
			Err:   errors.New("отрицательное число"),
		}
	case value == 0:
		return Result{
			Value: value,
			Err:   errors.New("число не может быть 0"),
		}
	default:
		return Result{
			Value: value * 2,
		}
	}
}
