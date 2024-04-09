package com.stos.worker;
import com.google.protobuf.InvalidProtocolBufferException;
import com.stos.worker.proto.Messages;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import java.time.LocalTime;
import java.util.Random;

@Component
public class Worker {

    private static final Random random = new Random();
    private static final String TASK_QUEUE_NAME = "ev_tasks";
    private static final String TASK_EXCHANGE_NAME = "stos";
    private static final String TASK_ROUTING_KEY = "ev_files";
    private static final String LOGGER_ROUTING_KEY = "log";
    private static final String FILE_SERVER_ROUTING_KEY = "files_server";

    private final Helpers helpers = new Helpers();

    @Autowired
    private RabbitTemplate rabbitTemplate;


    public void sendTaskMessage(byte[] message) {
        rabbitTemplate.convertAndSend(TASK_EXCHANGE_NAME, TASK_ROUTING_KEY, message);
        System.out.println("Task message sent.");
    }

    public void sendLogs(byte[] logs){
        rabbitTemplate.convertAndSend(TASK_EXCHANGE_NAME, LOGGER_ROUTING_KEY, logs);
        System.out.println("Log sent.");
    }

    public void logEvent(String content, Messages.LogType logType){
        Messages.LogEvent logEvent = Messages.LogEvent.newBuilder()
                .setTime(LocalTime.now().toString())
                .setType(logType)
                .setContent(content)
                .build();
        sendLogs(logEvent.toByteArray());
    }


    public void receiveMessageRabbit() throws InvalidProtocolBufferException {
        Message message = rabbitTemplate.receive(TASK_QUEUE_NAME);
        if (message != null) {
            Messages.TaskDispatch task = helpers.deserializeTaskDispatch(message.getBody());
            processTask(task);
            sendTaskMessage(task.toByteArray());
        }
    }

    public void processTask(Messages.TaskDispatch task) {
        System.out.println("Processing task:");
        System.out.println(task);
        processFiles();
        try {
            Thread.sleep(100);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        sendFiles();
        logEvent("log from worker!!!", Messages.LogType.forNumber(random.nextInt(3)));
        System.out.println("Task processed.");
    }

    public void workerTest() throws InvalidProtocolBufferException {
        while(true){
            receiveMessageRabbit();
        }
    }

    public void sendFiles(){
        //todo
    }

    public void processFiles(){
        //todo
    }
}