package rabbit

import (
	"fmt"
	"logger/stos-messages"
	"logger/utils"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/theritikchoure/logx"
	"google.golang.org/protobuf/proto"
)

var rabbit_conn amqp.Connection;

func InitRabbitConnection(retries int, timeout int) {
	rabit_addr, ok := os.LookupEnv("RABBIT_ADDR");
	if !ok {
		rabit_addr = "localhost"
	}
	connection_string := fmt.Sprintf("amqp://guest:guest@%s:5672/", rabit_addr);
	for i := 1; i<=retries; i++ {
		try := fmt.Sprintf("Trying to connect to rmq instance, try number: %d", i);
                logx.LogWithLevelAndTimestamp(try, "INFO", logx.FGWHITE, logx.BGMAGENTA)
		conn, err := amqp.Dial(connection_string);
		if err == nil {
			rabbit_conn = *conn;
			defer conn.Close()
			logx.LogWithLevelAndTimestamp("Connected successfully", "INFO", logx.FGWHITE, logx.BGMAGENTA)
			handleRabbitConnection()
		} else {
			logx.LogWithLevelAndTimestamp("Could not connect to rmq instance", "WARNING", logx.FGYELLOW, logx.BGMAGENTA)
			time.Sleep(time.Duration(timeout * int(time.Second)))
		}
	}
}

func handleRabbitConnection() {
	channel, err := rabbit_conn.Channel();
	defer channel.Close();
	utils.PanicOnError(err, "Error while creating rmq channel")
	queue, err := channel.QueueDeclare("log", false, false, false, false, nil);
	utils.PanicOnError(err, "Error while declaring queue");
	err = channel.ExchangeDeclare("stos", "direct", false, false, false, false, nil);
	utils.PanicOnError(err, "Error while declaring exchange");
	err = channel.QueueBind("log", "log", "stos", false, nil);
	utils.PanicOnError(err, "Error while bindind queue");

	var forever chan struct {};

	msgs, err := channel.Consume(
		queue.Name,
		"logger",
		true, // auto ack
		true, //exclusive
		false,
		false,
		nil,
	);
	utils.PanicOnError(err, "Error while starting consumer")
	
	go func() {
		for message := range msgs {
			go logMessage(message.Body)
		}
	}()
	logx.LogWithLevelAndTimestamp("Listening for messages", "INFO", logx.FGWHITE, logx.BGMAGENTA)
	<-forever
}

func logMessage(data []byte) error {
	event := &stos_messages.LogEvent{};
	if err := proto.Unmarshal(data, event); err != nil {
		return err
	}
	outs := fmt.Sprintf("{%s} @ %s", event.Content, event.Time)
	log_type := utils.EventToLoggerLevel(event.Type);
        logx.LogWithLevel(outs, log_type);
	return nil
}
