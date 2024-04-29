package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/KaranJagtiani/go-logstash"
	"github.com/google/uuid"
)

var tasks chan string;
var logger *log.Logstash;

type QueueError struct {}

func (q *QueueError) Error() string {
	return "Error while getting data from queue"
}

func getId() string{
	id, err := uuid.NewRandom();
	if err != nil {
		payload := map[string]interface{}{
			"message": "Error while generating uuid",
			"error":   true,
		}
		logger.Error(payload)
		id = uuid.Max;
	}
	return id.String()
}

func registerTask() {
	id := getId();
	tasks <- id
	payload := map[string]interface{}{
		"message": fmt.Sprintf("Registered task with uuid: %s", id),
		"error":   false,
	}
	logger.Log(payload)
}

func consumeTask() (string, error){
	if tasks == nil {
                return "", &QueueError{}
	}
	select {
	case id := <-tasks: 
        {
		payload := map[string]interface{}{
			"message": fmt.Sprintf("Consumed task: %s", id),
			"error":   false,
		}
		logger.Log(payload);
		return id, nil;
	}
	case <-time.After(time.Duration(3 * time.Second)): 
	{
		payload := map[string]interface{}{
			"message": "No tasks in queue, timed out",
			"error":   false,
		}
		logger.Log(payload)
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
		payload := map[string]interface{}{
			"message": "Received invalid request from worker",
			"error":   false,
		}
		logger.Warn(payload)
		w.WriteHeader(400)
		return
	}
	logger.LogString(fmt.Sprintf("Received results from worker for task: %s, processing time: %s miliseconds", body.ID, body.TIME))
}

func main() {
	logstashAddress, ok := os.LookupEnv("LOGSTASH_ADDR");
	if !ok {
		logstashAddress = "localhost"
	}
	logger = log.Init(logstashAddress, 5228, "tcp", 5);
	payload := map[string]interface{}{
		"message": "Starting evaluator",
		"error":   false,
	}
	logger.Info(payload)
	tasks = make(chan string, 10000)
	http.HandleFunc("/generate", handleTasks)
	http.HandleFunc("/consume", handleConsume)
	http.HandleFunc("/complete", handleComplete)
	payload = map[string]interface{}{
		"message": "Starting web server",
		"error":   false,
	}
	logger.Info(payload)
	http.ListenAndServe(":2137", nil)
}
