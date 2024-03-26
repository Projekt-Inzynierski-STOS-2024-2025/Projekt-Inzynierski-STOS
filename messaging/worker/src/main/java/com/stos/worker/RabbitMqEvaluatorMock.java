package com.stos.worker;
import com.google.protobuf.InvalidProtocolBufferException;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.amqp.core.Message;
@Component
public class RabbitMqEvaluatorMock {

    private static final String QUEUE_NAME = "evaluator_1";
    private static final String EXCHANGE_NAME = "evaluator_to_workers";
    private static final String ROUTING_KEY = "workers_queue";

    private final Helpers helpers = new Helpers();

    private final RabbitTemplate rabbitTemplate;

    @Autowired
    public RabbitMqEvaluatorMock(RabbitTemplate rabbitTemplate) {
        this.rabbitTemplate = rabbitTemplate;
    }

    public void sendMessage(byte[] message) {
        rabbitTemplate.convertAndSend(EXCHANGE_NAME, ROUTING_KEY, message);
    }

    public void sendGeneratedMessage() {
        this.sendMessage(helpers.generateTaskDispatch());
    }

    public String receiveMessage() throws InvalidProtocolBufferException {
        Message message = rabbitTemplate.receive(QUEUE_NAME);
        if (message != null) {
            return String.valueOf(helpers.deserializeTaskDispatch(message.getBody()));
        } else {
            return null;
        }
    }
}