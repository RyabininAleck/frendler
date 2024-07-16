package tasks

type Task interface {
	Run()
	Stop()
	execute()
}
