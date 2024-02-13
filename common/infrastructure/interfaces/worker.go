package interfaces

type Worker interface {
	start() error
	stop() error
}
