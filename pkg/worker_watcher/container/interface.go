package container

import (
	"github.com/spiral/roadrunner/v2/pkg/worker"
)

// Vector interface represents vector container
type Vector interface {
	// Enqueue used to put worker to the vector
	Enqueue(worker.BaseProcess)
	// Dequeue used to get worker from the vector
	Dequeue() (worker.BaseProcess, bool)
	// Destroy used to stop releasing the workers
	Destroy()
}
