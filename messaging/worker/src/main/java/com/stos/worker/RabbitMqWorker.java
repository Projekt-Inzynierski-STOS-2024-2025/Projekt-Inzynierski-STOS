package com.stos.worker;
import com.google.protobuf.InvalidProtocolBufferException;
import com.stos.worker.proto.Messages;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.Scanner;

@Component
public class RabbitMqWorker {

    private static final String QUEUE_NAME = "workers_queue";
    private static final String EXCHANGE_NAME = "worker_to_evaluator";
    private static final String ROUTING_KEY = "evaluator_1";

    private final Helpers helpers = new Helpers();

    @Autowired
    private RabbitTemplate rabbitTemplate;


    public void sendMessage(byte[] message) {
        rabbitTemplate.convertAndSend(EXCHANGE_NAME, ROUTING_KEY, message);
    }

    public void sendGeneratedMessage() {
        this.sendMessage(helpers.generateTaskDispatch());
    }

    public String receiveMessage() throws InvalidProtocolBufferException {
        Message message = rabbitTemplate.receive(QUEUE_NAME);
        if (message != null) {
            processTask(message);
            return "Task processed and sent back";
        } else {
            return null;
        }
    }

    public void processTask(Message message) throws InvalidProtocolBufferException {
        System.out.println("Processing task:");
        Messages.TaskDispatch task = helpers.deserializeTaskDispatch(message.getBody());
        System.out.println(task);
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("Task processed.");
        sendMessage(task.toByteArray());
    }

}