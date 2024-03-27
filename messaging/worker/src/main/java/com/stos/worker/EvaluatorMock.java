package com.stos.worker;
import com.google.protobuf.InvalidProtocolBufferException;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.common.serialization.ByteArrayDeserializer;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.amqp.core.Message;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;

import java.time.Duration;
import java.util.Collections;
import java.util.Properties;
@Component
public class EvaluatorMock {

    private static final String QUEUE_NAME = "evaluator_1";
    private static final String EXCHANGE_NAME = "evaluator_worker";
    private static final String ROUTING_KEY = "workers_queue";

    private final Helpers helpers = new Helpers();

    private final RabbitTemplate rabbitTemplate;


    @Autowired
    public EvaluatorMock(RabbitTemplate rabbitTemplate) {
        this.rabbitTemplate = rabbitTemplate;
    }

    public void sendMessageRabbit(byte[] message) {
        rabbitTemplate.convertAndSend(EXCHANGE_NAME, ROUTING_KEY, message);
    }

    public void sendGeneratedMessageRabbit() {
        this.sendMessageRabbit(helpers.generateTaskDispatch());
    }

    public String receiveMessageRabbit() throws InvalidProtocolBufferException {
        Message message = rabbitTemplate.receive(QUEUE_NAME);
        if (message != null) {
            return String.valueOf(helpers.deserializeTaskDispatch(message.getBody()));
        } else {
            return null;
        }
    }

    public void readFromKafkaTopic() {
        Properties props = new Properties();
        props.put(ConsumerConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:9092");
        props.put(ConsumerConfig.GROUP_ID_CONFIG, "worker-group");
        props.put(ConsumerConfig.KEY_DESERIALIZER_CLASS_CONFIG, ByteArrayDeserializer.class.getName());
        props.put(ConsumerConfig.VALUE_DESERIALIZER_CLASS_CONFIG, ByteArrayDeserializer.class.getName());

        try (KafkaConsumer<String, byte[]> consumer = new KafkaConsumer<>(props)) {
            consumer.subscribe(Collections.singletonList("evaluator_1"));

            while (true) {
                ConsumerRecords<String, byte[]> records = consumer.poll(Duration.ofMillis(100));
                for (ConsumerRecord<String, byte[]> record : records) {
                    System.out.println(helpers.deserializeTaskDispatch(record.value()));
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}