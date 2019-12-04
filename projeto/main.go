package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	. "arrebol/aux"
	. "arrebol/types"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	logrus "github.com/sirupsen/logrus"
)

var queues = make(map[string]QueueSpec)

func HomeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Arrebol is up and running.")
}

func CreateQueue(w http.ResponseWriter, r *http.Request) {
	var newQueue QueueSpec
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := ErrorMessage{"Bad Request"}
		json.NewEncoder(w).Encode(message)
	}

	var newID UUIDString
	id := GenerateID()
	newID.ID = id.String()
	newQueue.ID = id

	json.Unmarshal(reqBody, &newQueue)

	newQueue.Init()
	newQueue.MockWorkers()
	newQueue.Dispatch()
	queues[newID.ID] = newQueue

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newID)
}

func ListQueues(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	queueList := make([]QueueObj, 0)
	for _, q := range queues {
		qi := RecoverQueue(q)
		qi.Jobs = make([]JobSpec, 0)
		qi.Workers = make([]WorkerObj, 0)
		queueList = append(queueList, qi)
	}
	json.NewEncoder(w).Encode(queueList)
}

func GetQueue(w http.ResponseWriter, r *http.Request) {
	queueID := mux.Vars(r)["queue_id"]
	q := FindQueueByID(queueID, queues)
	if q == nil {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}
	queue := RecoverQueue(*q)
	json.NewEncoder(w).Encode(queue)
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
	queueID := mux.Vars(r)["queue_id"]
	queue := FindQueueByID(queueID, queues)

	if queue == nil {
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(message)
		return
	}

	var newWorker WorkerSpec
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		message := ErrorMessage{"Bad Request"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message)
		return
	}

	var newID UUIDString
	id := GenerateID()
	newID.ID = id.String()
	newWorker.ID = id

	json.Unmarshal(reqBody, &newWorker)
	queue.Workers[newID.ID] = newWorker

	queue.FreeWorkers <- newID.ID

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newID)
}

func ListWorkers(w http.ResponseWriter, r *http.Request) {
	queueID := mux.Vars(r)["queue_id"]
	queue := FindQueueByID(queueID, queues)

	if queue == nil {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}

	workers := TransformWorkers(queue.Workers)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(workers)
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queueID := vars["queue_id"]
	queue := FindQueueByID(queueID, queues)

	if queue == nil {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}

	workerID := vars["worker_id"]
	_, ok := queue.Workers[workerID]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The worker with id " + workerID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}

	delete(queue.Workers, workerID)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("")
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
	queueID := mux.Vars(r)["queue_id"]
	queue := FindQueueByID(queueID, queues)

	if queue == nil {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}

	var newJob InputJob
	reqBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		message := ErrorMessage{"Bad Request"}
		json.NewEncoder(w).Encode(message)
	}

	var newID UUIDString
	id := GenerateID()
	newID.ID = id.String()
	newJob.ID = id

	json.Unmarshal(reqBody, &newJob)
	job := TransformInputJob(newJob)
	queue.ToDoJobs[newID.ID] = job
	queue.JobsPending <- newID.ID

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newID)
}

func ListJobs(w http.ResponseWriter, r *http.Request) {
	queueID := mux.Vars(r)["queue_id"]
	queue := FindQueueByID(queueID, queues)
	q := RecoverQueue(*queue)

	if queue == nil {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}

	label := r.FormValue("label")
	state := r.FormValue("state")

	jobs := TransformJobs(q.Jobs)
	jobs = FilterJobsByLabel(jobs, label)
	jobs = FilterJobsByState(jobs, state)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jobs)
}

func GetJob(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queueID := vars["queue_id"]
	queue := FindQueueByID(queueID, queues)
	q := RecoverQueue(*queue)

	if queue == nil {
		w.WriteHeader(http.StatusNotFound)
		message := ErrorMessage{"The queue with id " + queueID + " was not found."}
		json.NewEncoder(w).Encode(message)
		return
	}

	jobID := vars["job_id"]
	jobUUID := uuid.FromStringOrNil(jobID)

	for _, job := range q.Jobs {
		if job.ID == jobUUID {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(job)
		}
	}
}

// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Println(r.Method+":", r.RequestURI)
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// func main() {
// 	q := QueueSpec{}
// 	q.Init()
// 	q.MockJobs()
// 	q.Dispatch()
// 	<-q.Killed
// }

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(logMiddleware)
	router.HandleFunc("/api/queues/{queue_id}/workers/{worker_id}", DeleteWorker).Methods("DELETE")
	router.HandleFunc("/api/queues/{queue_id}/workers", CreateWorker).Methods("POST")
	router.HandleFunc("/api/queues/{queue_id}/workers", ListWorkers).Methods("GET")
	router.HandleFunc("/api/queues/{queue_id}/jobs/{job_id}", GetJob).Methods("GET")
	router.HandleFunc("/api/queues/{queue_id}/jobs", ListJobs).Methods("GET").Queries("label", "{label_value}", "state", "{state_value}")
	router.HandleFunc("/api/queues/{queue_id}/jobs", ListJobs).Methods("GET").Queries("state", "{state_value}")
	router.HandleFunc("/api/queues/{queue_id}/jobs", ListJobs).Methods("GET").Queries("label", "{label_value}")
	router.HandleFunc("/api/queues/{queue_id}/jobs", ListJobs).Methods("GET")
	router.HandleFunc("/api/queues/{queue_id}/jobs", CreateJob).Methods("POST")
	router.HandleFunc("/api/queues/{queue_id}", GetQueue).Methods("GET")
	router.HandleFunc("/api/queues", CreateQueue).Methods("POST")
	router.HandleFunc("/api/queues", ListQueues).Methods("GET")
	router.HandleFunc("/", HomeFunc)

	log.Fatal(http.ListenAndServe(":8080", router))
}
