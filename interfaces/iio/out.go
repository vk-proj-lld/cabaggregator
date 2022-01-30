package iio

type IOout interface {
	Write(contents ...interface{})
}
