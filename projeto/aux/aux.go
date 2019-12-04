package aux

import . "arrebol/types"
import uuid "github.com/satori/go.uuid"
import "strings"
import "fmt"

func FindQueueByID(id string, queues map[string]QueueSpec) *QueueSpec {
	queue, ok := queues[id]
	if ok {
		return &queue
	}
	return nil
}

func FilterJobsByLabel(jobs []QueueJob, label string) []QueueJob {
	if label == "" {
		return jobs
	}
	filtered := make([]QueueJob, 0)
	for _, job := range jobs {
		if job.Label == label {
			filtered = append(filtered, job)
		}
	}
	return filtered
}

func FilterJobsByState(jobs []QueueJob, state string) []QueueJob {
	if state == "" {
		return jobs
	}
	filtered := make([]QueueJob, 0)
	for _, job := range jobs {
		if strings.ToLower(job.State) == state {
			filtered = append(filtered, job)
		}
	}
	return filtered
}

func GenerateID() uuid.UUID {
	id, error := uuid.NewV4()
	if error != nil {
		fmt.Println("Error generating the id")
	}
	return id
}

func GetJobState(job JobSpec) string {
	tasks := job.Tasks
	numQueued := 0
	for _, task := range tasks {
		for _, command := range task.Commands {
			if command.State == "RUNNING" {
				return "RUNNING"
			} else if command.State == "QUEUED" {
				numQueued++
			} else if command.State == "FAILED" {
				return "FAILED"
			}
		}
	}
	if numQueued > 0 {
		return "QUEUED"
	}
	return "FINISHED"
}

// TransformWorkers transforms WorkerSpec into WorkerObj
func TransformWorkers(workers []WorkerSpec) []WorkerObj {
	objs := make([]WorkerObj, 0)
	for _, w := range workers {
		worker := &WorkerObj{}
		worker.ID = w.ID.String()
		worker.Address = w.Address
		worker.PoolSize = w.PoolSize

		objs = append(objs, *worker)
	}
	return objs
}

// TransformJobs transforms Job into QueueJob
func TransformJobs(jobs []JobSpec) []QueueJob {
	objs := make([]QueueJob, 0)
	for _, j := range jobs {
		qjob := QueueJob{}
		qjob.ID = j.ID.String()
		qjob.Label = j.Label
		qjob.State = j.State
		objs = append(objs, qjob)
	}
	return objs
}

func TransformInputJob(input InputJob) JobSpec {
	job := JobSpec{}
	job.ID = input.ID
	job.Label = input.Label
	job.State = "QUEUED"
	tasks := make([]TaskSpec, 0)
	for _, ti := range input.Tasks {
		commands := make([]CommandSpec, 0)
		for _, ci := range ti.Commands {
			com := CommandSpec{}
			com.Command = ci
			com.State = "QUEUED"
			com.ExitCode = 0
			commands = append(commands, com)
		}
		task := TaskSpec{}
		task.Spec = ti.Spec
		task.Commands = commands
		tasks = append(tasks, task)
	}
	job.Tasks = tasks
	return job
}

func RecoverQueue(queue QueueSpec) QueueObj {
	q := QueueObj{}
	q.Jobs = MergeJobs(queue.JobsArrived, queue.JobsWaiting, queue.JobsRunning, queue.JobsFinished)
	q.Workers = TransformWorkers(MergeWorkers(queue.IdleWorkers, queue.BusyWorkers))
	q.ID = queue.ID
	q.Name = queue.Name
	q.PoolsSize = q.PoolsSize
	return q
}

func MergeJobs(arrived map[string]JobSpec, todo map[string]JobSpec,
	doing map[string]JobSpec, done map[string]JobSpec) []JobSpec {
	jobs := make([]JobSpec, 0)
	for _, j := range arrived {
		jobs = append(jobs, j)
	}
	for _, j := range todo {
		jobs = append(jobs, j)
	}
	for _, j := range doing {
		jobs = append(jobs, j)
	}
	for _, j := range done {
		jobs = append(jobs, j)
	}
	return jobs
}

func MergeWorkers(idle map[string]WorkerSpec, busy map[string]WorkerSpec) []WorkerSpec {
	workers := make([]WorkerSpec, 0)
	for _, w := range idle {
		workers = append(workers, w)
	}
	for _, w := range busy {
		workers = append(workers, w)
	}
	return workers
}
