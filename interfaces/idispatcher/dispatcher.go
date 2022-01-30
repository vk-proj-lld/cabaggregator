package idispatcher

type IDispatcher interface {
	AddDriver(ids ...string)
	Dispatch() error
}
