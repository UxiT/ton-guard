package common

type Request interface {
	Validate() error
}
