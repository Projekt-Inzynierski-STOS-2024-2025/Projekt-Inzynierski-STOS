package main

import (
	"fmt"
	"logger/stos-messages"
	"time"

	"google.golang.org/protobuf/proto"
)


func main() {
	fmt.Println("Starting logger");
	event := stos_messages.LogEvent{Content: "Some content", Time: time.Now().String(), Type: stos_messages.LogType_INFO};
	fmt.Println(event.Content)
	bin, err := proto.Marshal(&event);
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bin)
}
