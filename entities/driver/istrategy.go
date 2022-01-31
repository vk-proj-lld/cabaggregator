package driver

type IStrategy interface {
	Select() AckSignal
}
