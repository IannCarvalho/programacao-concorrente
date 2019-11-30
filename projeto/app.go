package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "io/ioutil"
)

// >>>>>>>>>>>>>>>>>>>>>>> STRUCTS

type QueueSpec struct {
    Id              string          `json:"label"`
    Name            string          `json:"name"`
    WaitingJobs     int             `json:"waiting_jobs"`
    WorkerPools     int             `json:"worker_pools"`
    PoolsSize       int             `json:"pools_size"`
    Jobs            []JobSpec       `json:"jobs,omitempty"`
    Workers         []WorkerSpec    `json:"workers,omitempty"`
}

type WorkerSpec struct {
    Id              string  `json:"id"`
    Address         string  `json:"address"`
    PoolSize        int     `json:"pool_size"`
}

type TaskSpec struct {
    Image           string          `json:"image"`
    Requirements    struct {
        DockerRequirements  string  `json:"DockerRequirements"`
    }                               `json:"requirements"`
}

type JobSpec struct {
    Label       string          `json:"label"`
    Tasks       []struct {
        Id          string      `json:"id"`
        Spec        TaskSpec    `json:"spec"`
        Commands    []string    `json:"commands"`
    }                           `json:"tasks,omitempty"`
}

type Job struct {
    Id      string  `json:"id"`
    JobSpec JobSpec `json:"jobSpec"`
}

type ErrorMessage struct {
    Message     string      `json:"message"`
}

// <<<<<<<<<<<<<<<<<<<<<<< STRUCTS
// >>>>>>>>>>>>>>>>>>>>>>> AUX FUNCS

func FindQueueById(id string) *QueueSpec {
    for _, queue := range queues {
        if queue.Id == id {
            return &queue
        }
    }
    return &QueueSpec{}
}

// <<<<<<<<<<<<<<<<<<<<<<< AUX FUNCS

type allQueues []QueueSpec
var queues = allQueues{
    {
        Id: "a3f77cca-96bb-4e7a-b746-e8960f779747",
        Name: "awesome_name",
        WaitingJobs: 0,
        WorkerPools: 0,
        PoolsSize: 0,
    },
}

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

    json.Unmarshal(reqBody, &newQueue)
    queues = append(queues, newQueue)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newQueue)
}

func ListQueues(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    queueList := queues
    for _, qi := range queueList {
        qi.Workers = make([]WorkerSpec, 0)
        qi.Jobs =  make([]JobSpec, 0)
    }
    json.NewEncoder(w).Encode(queueList)
}

func GetQueue(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    q := FindQueueById(queueId)
    if q == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
    q.Workers = make([]WorkerSpec, 0)
    q.Jobs = make([]JobSpec, 0)
    json.NewEncoder(w).Encode(q)
}

func CreateWorker(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }

    var newWorker WorkerSpec
    reqBody, error := ioutil.ReadAll(r.Body)
    if error != nil {
        w.WriteHeader(http.StatusBadRequest)
        message := ErrorMessage{"Bad Request"}
        json.NewEncoder(w).Encode(message)
    }

    json.Unmarshal(reqBody, &newWorker)
    queue.Workers = append(queue.Workers, newWorker)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newWorker)
}

func ListWorkers(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }

    workers := queue.Workers
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(workers)
}

func DeleteWorker(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
}

func CreateJob(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
}

func ListJobs(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
}

func GetJob(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
}

func ListLabelJobs(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
}

func ListStateJobs(w http.ResponseWriter, r *http.Request) {
    queueId := mux.Vars(r)["queue_id"]
    queue := FindQueueById(queueId)

    if queue == nil {
        w.WriteHeader(http.StatusNotFound)
        message := ErrorMessage{"The queue with id " + queueId + " was not found."}
        json.NewEncoder(w).Encode(message)
    }
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", HomeFunc)
    router.HandleFunc("/api/queues", ListQueues).Methods("GET")
    router.HandleFunc("/api/queues", CreateQueue).Methods("POST")
    router.HandleFunc("/api/queues/{queue_id}", GetQueue).Methods("GET")
    router.HandleFunc("/api/queues/{queue_id}/workers", ListWorkers).Methods("GET")
    router.HandleFunc("/api/queues/{queue_id}/workers", CreateWorker).Methods("POST")
    router.HandleFunc("/api/queues/{queue_id}/workers/{worker_id}", DeleteWorker).Methods("DELETE")

    router.HandleFunc("/api/queues/{queue_id}/jobs", CreateJob).Methods("POST")
    router.HandleFunc("/api/queues/{queue_id}/jobs", ListJobs).Methods("GET")
    router.HandleFunc("/api/queues/{queue_id}/jobs", ListLabelJobs).Methods("GET").Queries("label", "{label_value}")
    router.HandleFunc("/api/queues/{queue_id}/jobs", ListStateJobs).Methods("GET").Queries("state", "{state_value}")
    router.HandleFunc("/api/queues/{queue_id}/jobs/{job_id}", GetJob).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
