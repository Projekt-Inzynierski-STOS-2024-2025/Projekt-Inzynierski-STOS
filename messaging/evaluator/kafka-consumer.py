from confluent_kafka import Consumer

KAFKA_HOST: str = 'localhost:9092'
TOPIC_NAME: str = 'kafka-test-topic'
GROUP_ID: str = 'STOS'


def consume_messages():
    message_count: int = 0
    conf = {
        'bootstrap.servers': 'localhost:9092',
        'group.id': GROUP_ID,
        'auto.offset.reset': 'earliest'
    }

    consumer = Consumer(conf)
    consumer.subscribe([TOPIC_NAME])
    print("Subscribed")

    while True:
        msg = consumer.poll(timeout=1.0)
        if msg is not None:
            message_count += 1
            print(f'Received message no. {message_count}: {msg.value().decode("utf-8")}')


if __name__ == '__main__':
    print("Kafka consumer - start")
    consume_messages()
