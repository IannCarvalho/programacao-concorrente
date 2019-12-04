package types

import (
	uuid "github.com/satori/go.uuid"
	logrus "github.com/sirupsen/logrus"
	"math"
	"sync"
	"time"
)

// UUIDString is used in return in the http response
type UUIDString struct {
	ID string `json:"id"`
}

// QueueSpec is the Queue itself
type QueueSpec struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	WaitingJobs int       `json:"waiting_jobs"`
	WorkerPools int       `json:"worker_pools"`
	PoolsSize   int       `json:"pools_size"`

	ArrivedLock  sync.Mutex
	JobsArrived  map[string]JobSpec
	JobsWaiting  map[string]JobSpec
	JobsRunning  map[string]JobSpec
	JobsFinished map[string]JobSpec

	IdleWorkers map[string]WorkerSpec `json:"workers,omitempty"`
	BusyWorkers map[string]WorkerSpec

	NotifyWorkerIsIdle chan string
	NotifyNewJob       chan string
	Helper             chan string

	Killed chan bool
}

type QueueObj struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	WaitingJobs int         `json:"waiting_jobs"`
	WorkerPools int         `json:"worker_pools"`
	PoolsSize   int         `json:"pools_size"`
	Jobs        []JobSpec   `json:"jobs,omitempty"`
	Workers     []WorkerObj `json:"workers,omitempty"`
}

// WorkerSpec is the Worker used internally
type WorkerSpec struct {
	ID       uuid.UUID `json:"id"`
	Address  string    `json:"address"`
	PoolSize int       `json:"pool_size"`
}

// WorkerObj is the Worker we use as output in the http response
type WorkerObj struct {
	ID       string `json:"id"`
	Address  string `json:"address"`
	PoolSize int    `json:"pool_size"`
}

// ImageSpec is part of a TaskSpec
type ImageSpec struct {
	Image        string `json:"image"`
	Requirements struct {
		DockerRequirements string `json:"DockerRequirements"`
	} `json:"requirements"`
}

// TaskSpec is part of a JobSpec
type TaskSpec struct {
	ID       string        `json:"id"`
	Spec     ImageSpec     `json:"spec"`
	Commands []CommandSpec `json:"commands"`
}

// JobSpec is the Job used internally
type JobSpec struct {
	ID    uuid.UUID  `json:"id"`
	Label string     `json:"label,omitempty"`
	State string     `json:"state,omitempty"`
	Tasks []TaskSpec `json:"tasks,omitempty"`
}

// QueueJob is the Job used as output in the http response
type QueueJob struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	State string `json:"state"`
}

// InputJob is the Job given as input in the http request
type InputJob struct {
	ID    uuid.UUID `json:"id"`
	Label string    `json:"label"`
	Tasks []struct {
		ID       string    `json:"id"`
		Spec     ImageSpec `json:"spec"`
		Commands []string  `json:"commands"`
	} `json:"tasks"`
}

// CommandSpec is the command used internally
type CommandSpec struct {
	Command  string `json:"command"`
	State    string `json:"state"`
	ExitCode int    `json:"exit_code"`
}

// ErrorMessage is used in the http reponse for errors
type ErrorMessage struct {
	Message string `json:"message"`
}

type Allocation struct {
	worker WorkerSpec
	job    JobSpec
}

func (q *QueueSpec) MockWorkers() {
	rng := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, i := range rng {
		nw := WorkerSpec{}
		nw.Address = i
		nw.ID, _ = uuid.NewV4()
		nw.PoolSize = 0
		q.IdleWorkers[nw.ID.String()] = nw
	}
}

func (q *QueueSpec) MockJobs() {

	rng := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, i := range rng {
		command := CommandSpec{}
		commands := make([]CommandSpec, 0)
		commands = append(commands, command)

		task := TaskSpec{}
		task.ID = "Teste"
		task.Commands = commands
		tasks := make([]TaskSpec, 0)
		tasks = append(tasks, task)

		job := JobSpec{}
		job.ID, _ = uuid.NewV4()
		job.Label = i
		job.Tasks = tasks

		q.JobsArrived[job.ID.String()] = job
		q.NotifyNewJob <- job.ID.String()
	}
}

func (q *QueueSpec) Init() {
	q.IdleWorkers = make(map[string]WorkerSpec)
	q.BusyWorkers = make(map[string]WorkerSpec)

	q.JobsArrived = make(map[string]JobSpec, 10)
	q.JobsWaiting = make(map[string]JobSpec, 10)
	q.JobsRunning = make(map[string]JobSpec, 10)
	q.JobsFinished = make(map[string]JobSpec, 10)

	q.NotifyNewJob = make(chan string, 100)
	q.NotifyWorkerIsIdle = make(chan string, 100)
	q.Helper = make(chan string, 100)
	q.Killed = make(chan bool, 1)
}

func (q *QueueSpec) Dispatch() {
	logrus.Println("Queue", q.ID, "started dispatching.")
	go func() {
		for {
			select {

			case workerID := <-q.NotifyWorkerIsIdle:
				// Since channel does not accept tuples,
				// we are using a helper channel to store
				// the finished job id. When the worker is
				// idle, both channels will be notified
				jobID := <-q.Helper

				// MOVE JOB FROM RUNNING TO FINISHED
				job := q.JobsRunning[jobID]
				delete(q.JobsRunning, jobID)
				q.JobsFinished[jobID] = job

				// MOVE WORKER TO IDLE
				worker := q.BusyWorkers[workerID]
				delete(q.BusyWorkers, workerID)
				q.IdleWorkers[workerID] = worker

				// TRY MATCHING
				matches := TryAllocating(q.IdleWorkers, q.JobsWaiting)
				Allocate(q, matches)

			case jobID := <-q.NotifyNewJob:
				// MOVE NEW TO WAITING
				q.ArrivedLock.Lock()
				job := q.JobsArrived[jobID]
				delete(q.JobsArrived, jobID)
				q.ArrivedLock.Unlock()
				q.JobsWaiting[jobID] = job

				// TRY ALLOC (REPEAT)
				matches := TryAllocating(q.IdleWorkers, q.JobsWaiting)
				Allocate(q, matches)

			}
		}
	}()
}

func (w *WorkerSpec) Consume(j *JobSpec, freeWorkers chan string, jobsDone chan string) {
	logrus.Println("Worker", w.ID, "is working on job", j.ID)
	time.Sleep(100 * time.Millisecond)
	// j.State = "RUNNING"
	// tasks := j.Tasks
	// for m, task := range tasks {
	// 	for n, command := range task.Commands {
	// 		if command.State == "QUEUED" {
	// 			command.State = "RUNNING"
	// 			time.Sleep(100 * time.Millisecond)
	// 			command.State = "FINISHED"
	// 			command.ExitCode = 0
	// 		}
	// 		task.Commands[n] = command
	// 	}
	// 	tasks[m] = task
	// }
	// j.Tasks = tasks
	// j.State = "FINISHED"
	freeWorkers <- w.ID.String()
	jobsDone <- j.ID.String()
	logrus.Println("Worker", w.ID, "became idle")
}

func TryAllocating(idleWorkers map[string]WorkerSpec, pendingJobs map[string]JobSpec) []Allocation {
	allocations := make([]Allocation, 0)

	v := int(math.Min(float64(len(idleWorkers)), float64(len(pendingJobs))))

	workerKeys := make([]string, 0, v)
	i := 0
	for k := range idleWorkers {
		if i >= v {
			break
		}
		workerKeys = append(workerKeys, k)
		i++
	}

	jobKeys := make([]string, 0, v)
	j := 0
	for k := range pendingJobs {
		if j >= v {
			break
		}
		jobKeys = append(jobKeys, k)
		j++
	}

	for i := 0; i < v; i++ {
		worker := idleWorkers[workerKeys[i]]
		job := pendingJobs[jobKeys[i]]
		aux := Allocation{worker, job}
		allocations = append(allocations, aux)

	}
	return allocations
}

func Allocate(q *QueueSpec, matches []Allocation) {
	// IF MATCHING
	for _, wj := range matches {
		// MOVE JOB WAITING TO RUNNING
		job := wj.job
		delete(q.JobsWaiting, job.ID.String())
		q.JobsRunning[job.ID.String()] = job

		// MOVE WORKER IDLE TO BUSY
		worker := wj.worker
		delete(q.IdleWorkers, worker.ID.String())
		q.BusyWorkers[worker.ID.String()] = worker
	}

	for _, wj := range matches {
		worker := wj.worker
		job := wj.job
		go worker.Consume(&job, q.NotifyWorkerIsIdle, q.Helper)
	}
}
