from confluent_kafka import Producer
import sys
import time

KAFKA_HOST: str = 'localhost:9092'
TOPIC_NAME: str = 'kafka-test-topic'
GROUP_ID: str = 'STOS'
MAX_QUEUE_AMOUNT: int = 100_000


def produce_messages(amount: int):
    conf = {
        'bootstrap.servers': KAFKA_HOST
    }
    producer = Producer(conf)

    for i in range(amount):
        producer.produce(TOPIC_NAME, value=f'Message {i}')
    producer.flush()


if __name__ == '__main__':
    print("Kafka producer - start")
    amount: int = int(sys.argv[1])
    start_time = time.time()
    produce_messages(amount)

    print(f"Sent {amount} messages in {time.time() - start_time} seconds")
