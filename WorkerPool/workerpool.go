package workerpool

import "log"

type WorkerPool interface {
	Run()
	AddTask() (task func())
}

type T = interface{}

type workerPool struct {
	maxworker   int
	queuedTaskC chan func()
}

func NewWorkerPool(maxWorker int) *workerPool {
	wp := &workerPool{
		maxworker:   maxWorker,
		queuedTaskC: make(chan func()),
	}
	return wp
}

func (wp *workerPool) Run() {
	wp.run()
}

func (wp *workerPool) AddTask(task func()) {
	wp.queuedTaskC <- task
}

func (wp *workerPool) run() {
	for i := 0; i < wp.maxworker; i++ {
		wID := i + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)

		go func(workerId int) {
			for task := range wp.queuedTaskC {
				log.Printf("[WorkerPool] Worker %d starts processing task", wID)
				task()
				log.Printf("[WorkerPool] Worker %d finished processing task", wID)
			}
		}(wID)
	}
}
