package com.stos.worker;
import com.google.protobuf.InvalidProtocolBufferException;
import com.stos.worker.proto.Messages;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class Worker {

    private static final String TASK_QUEUE_NAME = "ev_tasks";
    private static final String TASK_EXCHANGE_NAME = "stos";
    private static final String TASK_ROUTING_KEY = "ev_files";
    private static final String LOGGER_ROUTING_KEY = "logger";

    private final Helpers helpers = new Helpers();

    @Autowired
    private RabbitTemplate rabbitTemplate;


    public void sendTaskMessage(byte[] message) {
        rabbitTemplate.convertAndSend(TASK_EXCHANGE_NAME, TASK_ROUTING_KEY, message);
    }


    public void receiveMessageRabbit() throws InvalidProtocolBufferException {
        Message message = rabbitTemplate.receive(TASK_QUEUE_NAME);
        if (message != null) {
            Messages.TaskDispatch task = helpers.deserializeTaskDispatch(message.getBody());
            processTask(task);
            sendTaskMessage(task.toByteArray());
            System.out.println("Task processed and sent back");
        }
    }

    public void processTask(Messages.TaskDispatch task) {
        System.out.println("Processing task:");
        System.out.println(task);
        requestFiles();
        processFiles();
        sendFiles();
        try {
            Thread.sleep(20);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("Task processed.");
    }

    public void workerTest() throws InvalidProtocolBufferException {
        while(true){
            receiveMessageRabbit();
        }
    }

    public void requestFiles(){
        //todo
    }

    public void sendFiles(){
        //todo
    }

    public void processFiles(){
        //todo
    }

    public void sendLogs(){
        //todo
    }
}