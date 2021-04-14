package utils

type Value interface {
	Len() int
}

type CallBackFunc func(key string, value Value)
