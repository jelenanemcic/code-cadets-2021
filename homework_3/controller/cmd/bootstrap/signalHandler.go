package bootstrap

import "github.com/jelenanemcic/code-cadets-2021/homework_3/controller/internal/tasks"

func SignalHandler() *tasks.SignalHandler {
	return tasks.NewSignalHandler()
}
