package com.stos.worker;
import com.google.protobuf.InvalidProtocolBufferException;
import com.stos.worker.proto.Messages;
import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;
import java.util.Properties;

@Component
public class Worker {

    private static final String QUEUE_NAME = "workers_queue";
    private static final String EXCHANGE_NAME = "evaluator_worker";
    private static final String ROUTING_KEY = "evaluator_1";

    private final Helpers helpers = new Helpers();

    @Autowired
    private RabbitTemplate rabbitTemplate;

    private KafkaProducer kafkaProducer;


    public void sendMessageRabbit(byte[] message) {
        rabbitTemplate.convertAndSend(EXCHANGE_NAME, ROUTING_KEY, message);
    }

    public void sendGeneratedMessageRabbit() {
        this.sendMessageRabbit(helpers.generateTaskDispatch());
    }


    public String receiveMessageRabbit() throws InvalidProtocolBufferException {
        Message message = rabbitTemplate.receive(QUEUE_NAME);
        if (message != null) {
            Messages.TaskDispatch task = helpers.deserializeTaskDispatch(message.getBody());
            processTask(task);
            sendMessageRabbit(task.toByteArray());
            return "Task processed and sent back";
        } else {
            return null;
        }
    }

    public void processTask(Messages.TaskDispatch task) throws InvalidProtocolBufferException {
        System.out.println("Processing task:");
        System.out.println(task);
        try {
            Thread.sleep(2000);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
        System.out.println("Task processed.");
    }

    public void sendGeneratedMessageKafka() {
        this.writeToKafkaTopic(helpers.generateTaskDispatch());
    }
    public void writeToKafkaTopic(byte[] message) {
        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");
        props.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        props.put("value.serializer", "org.apache.kafka.common.serialization.ByteArraySerializer");
        kafkaProducer =  new KafkaProducer<>(props);
        try {
            ProducerRecord<String, byte[]> record = new ProducerRecord<>("evaluator_1", message);
            kafkaProducer.send(record);
            kafkaProducer.flush();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}