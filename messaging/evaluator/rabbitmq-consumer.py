from pika import BlockingConnection
from pika import ConnectionParameters


QUEUE_NAME: str = 'rabbitmq-test-queue'
HOST_NAME: str = 'localhost'


def consume_messages():
    con = BlockingConnection(ConnectionParameters(host=HOST_NAME))
    channel = con.channel()
    channel.queue_declare(queue=QUEUE_NAME)

    channel.basic_consume(queue=QUEUE_NAME, on_message_callback=on_message)
    channel.start_consuming()


def on_message(channel, method, properties, body):
    print(body.decode())


if __name__ == '__main__':
    print("RabbitMQ consumer - start")
    consume_messages()
