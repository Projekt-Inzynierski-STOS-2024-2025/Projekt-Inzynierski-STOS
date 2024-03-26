from pika import BlockingConnection
from pika import ConnectionParameters
import time
import sys


QUEUE_NAME: str = 'rabbitmq-test-queue'
HOST_NAME: str = 'localhost'


def produce_messages(amount: int):
    con = BlockingConnection(ConnectionParameters(host=HOST_NAME))
    channel = con.channel()
    channel.queue_declare(queue=QUEUE_NAME)

    for i in range(amount):
        message: str = f'Message {i}'
        channel.basic_publish(body=message.encode(), routing_key=QUEUE_NAME, exchange='')

    con.close()


if __name__ == '__main__':
    print("RabbitMQ producer - start")
    amount: int = int(sys.argv[1])
    start_time = time.time()
    produce_messages(amount)

    print(f"Sent {amount} messages in {time.time() - start_time} seconds")
