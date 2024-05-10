package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"net"
	"github.com/google/uuid"
)

var (
	tasks chan string
	logger 	*Logger 
)

type QueueError struct {}

const (
    LogLevelDebug = iota
    LogLevelInfo
    LogLevelWarning
    LogLevelError
)

type GelfMessage struct {
    Host         string  `json:"host"`
    ShortMessage string  `json:"short_message"`
    Timestamp    float64 `json:"timestamp"`
    Level        int     `json:"level"`
	Duration 	 int     `json:_duration,omitempty`
	Error 		 bool	 `json:_error`
}

type Logger struct {
    graylogAddr string
    appName     string
}

type LogOptions struct {
	level	  int
    message   string
    duration  int
    isError   bool
}

func NewLogger(graylogAddr, appName string) *Logger {
    return &Logger{graylogAddr: graylogAddr, appName: appName}
}

func (l *Logger) Log(options LogOptions) {
    gelfMsg := GelfMessage{
        Host:         l.appName,
        ShortMessage: options.message,
        Timestamp:    float64(time.Now().UnixNano()) / 1e9,
        Level:        options.level,
		Error:		  options.isError,
    }

	if options.duration >= 0 {
        gelfMsg.Duration = options.duration
    }
	
    b, err := json.Marshal(gelfMsg)
    if err != nil {
        fmt.Println("Error encoding GELF message:", err)
        return
    }

    conn, err := net.Dial("udp", l.graylogAddr)
    if err != nil {
        fmt.Println("Error dialing Graylog server:", err)
        return
    }
    defer conn.Close()

    _, err = conn.Write(b)
    if err != nil {
        fmt.Println("Error sending GELF message:", err)
    }
}

func (q *QueueError) Error() string {
	return "Error while getting data from queue"
}

func getId() string{
	id, err := uuid.NewRandom();
	if err != nil {
		logger.Log(LogOptions{
			level: LogLevelError,
			message: "Error while generating uuid",
			duration: -1,
			isError: true,
		})
		id = uuid.Max;
	}
	return id.String()
}

func registerTask() {
	id := getId();
	tasks <- id
	logger.Log(LogOptions{
		level: LogLevelInfo,
		message: fmt.Sprintf("Registered task with uuid: %s", id),
		duration: -1,
		isError: false,
	})
}

func consumeTask() (string, error){
	if tasks == nil {
                return "", &QueueError{}
	}
	select {
	case id := <-tasks: 
        {
		logger.Log(LogOptions{
			level: LogLevelInfo,
			message: fmt.Sprintf("Consumed task: %s", id),
			duration: -1,
			isError: false,
		})
		return id, nil;
	}
	case <-time.After(time.Duration(3 * time.Second)): 
	{
		logger.Log(LogOptions{
			level: LogLevelInfo,
			message: "No tasks in queue, timed out",
			duration: -1,
			isError: false,
		})
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
		logger.Log(LogOptions{
			level: LogLevelWarning,
			message: "Received invalid request from worker",
			duration: -1,
			isError: false,
		})
		w.WriteHeader(400)
		return
	}
	time, err := strconv.Atoi(body.TIME);
	error := false;
	if err != nil {
		time = 0;
		error = true;

	}
	
	logger.Log(LogOptions{
		level: LogLevelInfo,
		message: fmt.Sprintf("Received worker result for task: %s", body.ID),
		duration: time,
		isError: error,
	})
}

func main() {
	appName := "evaluator"
	graylogAddress, ok := os.LookupEnv("GRAYLOG_ADDR")
	if !ok {
		graylogAddress = "localhost"
	}
	logger = NewLogger(graylogAddress, appName)

	logger.Log(LogOptions{
		level: LogLevelInfo,
		message: "Starting evaluator",
		duration: -1,
		isError: false,
	})

	tasks = make(chan string, 10000)
	http.HandleFunc("/generate", handleTasks)
	http.HandleFunc("/consume", handleConsume)
	http.HandleFunc("/complete", handleComplete)
	
	logger.Log(LogOptions{
		level: LogLevelInfo,
		message: "Starting web server",
		duration: -1,
		isError: false,
	})
	http.ListenAndServe(":2137", nil)
}
