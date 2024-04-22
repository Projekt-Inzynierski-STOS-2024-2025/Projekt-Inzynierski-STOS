package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var tasks chan string;

type QueueError struct {}

func (q *QueueError) Error() string {
	return "Error while getting data from queue"
}

func getId() string{
	id, err := uuid.NewRandom();
	if err != nil {
		log.Error("Error while generating uuid")
		id = uuid.Max;
	}
	return id.String()
}

func registerTask() {
	id := getId();
	tasks <- id
	log.Infof("Registered task with uuid: %s", id)
}

func consumeTask() (string, error){
	if tasks == nil {
                return "", &QueueError{}
	}
	select {
	case id := <-tasks: 
        {
		log.Infof("Consumed task: %s", id);
		return id, nil;
	}
	case <-time.After(time.Duration(3 * time.Second)): 
	{
		log.Warn("No tasks in queue, timed out")
		return "", &QueueError{}
	}
	}
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
        n := r.URL.Query().Get("n");
	taskAmount, err := strconv.Atoi(n);
	if err != nil {
		taskAmount = 1;
	}
	for range taskAmount {
		go registerTask()
	}
}

func handleConsume(w http.ResponseWriter, r *http.Request) {
	id, err := consumeTask();
	if err != nil {
		w.WriteHeader(400);
		return
	}
	w.WriteHeader(200);
	fmt.Fprintln(w, id);
}

type CompleteRequestDTO struct {
	ID string `json:"id"`
	TIME string `json:"time"`	
}

func handleComplete(w http.ResponseWriter, r *http.Request) {
        var body CompleteRequestDTO;

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&body);
	if err != nil {
		log.Warnf("Received invalid request from worker")
		w.WriteHeader(400)
		return
	}
	log.Infof("Received results from worker for task: %s, processing time: %s miliseconds", body.ID, body.TIME)
}

func main() {
	log.Info("Starting evaluator")
	tasks = make(chan string, 10000)
	http.HandleFunc("/generate", handleTasks)
	http.HandleFunc("/consume", handleConsume)
	http.HandleFunc("/complete", handleComplete)
	log.Info("Started web server")
	http.ListenAndServe(":2137", nil)
}
