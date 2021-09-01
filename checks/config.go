package checks

type Config interface {
	AddService() Check
}
