# Mock evaluator

Basically a http server with 3 endpoints

## GET /generate?n={tasks}
Generates n tasks with random uuid's and adds them to the internal queue

## GET /consume
Consumes a task from queue and returns its uuid

## POST /complete
``` json 
{
    "id": "some uuid"
}
```
Notifies the server that a task has been completed
