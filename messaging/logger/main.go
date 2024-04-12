package main

import (
	rmq "logger/rabbit"
	"github.com/theritikchoure/logx"
)


func main() {
	logx.LogWithLevelAndTimestamp("Starting logger", "INFO", logx.FGWHITE, logx.BGMAGENTA);
	rmq.InitRabbitConnection(10, 5);
}
