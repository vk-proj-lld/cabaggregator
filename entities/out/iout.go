package out

type IOout interface {
	Write(contents ...interface{})
}
