from pika import BlockingConnection
from pika import ConnectionParameters
from flask import Flask, request
import random
import threading
import time
import messages_pb2


HOST_NAME: str = 'rabbitmq'
EXCHANGE_NAME: str = 'stos'
QUEUE_TO_WORKER_NAME: str = 'ev_tasks'
QUEUE_TO_FILES_NAME: str = 'files'
QUEUE_FROM_WORKER_NAME: str = 'ev_files'
FILESERVER_ROUTING_KEY = 'file_server'


sender_connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
sender_channel = sender_connection.channel()
receiver_connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
receiver_channel = receiver_connection.channel()


def setup_channel():
    sender_channel.exchange_declare(exchange=EXCHANGE_NAME, exchange_type='direct')
    sender_channel.queue_declare(queue=QUEUE_TO_WORKER_NAME)
    sender_channel.queue_declare(queue=QUEUE_TO_FILES_NAME)
    sender_channel.queue_bind(exchange=EXCHANGE_NAME, queue=QUEUE_TO_FILES_NAME, routing_key=FILESERVER_ROUTING_KEY)
    sender_channel.queue_bind(exchange=EXCHANGE_NAME, queue=QUEUE_TO_WORKER_NAME, routing_key=QUEUE_TO_WORKER_NAME)

    receiver_channel.queue_declare(queue=QUEUE_FROM_WORKER_NAME)
    receiver_channel.queue_bind(exchange=EXCHANGE_NAME, queue=QUEUE_FROM_WORKER_NAME, routing_key=QUEUE_FROM_WORKER_NAME)


def resolve_message(ch, parameters, method, body):
    message = messages_pb2.TaskDispatch()
    message.ParseFromString(body)
    time.sleep(random.random() * 2)

    print("---Evaluator received message---")
    print(message.student_id)
    print(message.task_id)
    print(message.files_hash)
    print(message.data)
    print("---Message end---")


def start_consuming():
    receiver_channel.basic_consume(queue=QUEUE_FROM_WORKER_NAME, on_message_callback=resolve_message, auto_ack=True)
    print("Evaluator consumer - start")
    receiver_channel.start_consuming()


def send_message(task_id: str, student_id: str, files_hash: bytes, files: list):
    files_message_content = []
    for filename in files:
        file_message = messages_pb2.File()
        file_message.name = filename
        file_message.data = open(filename, "rb").read()
        files_message_content.append(file_message)

    files_message = messages_pb2.Files()
    files_message.uuid = "Test"
    files_message.content.extend(files_message_content)

    message = messages_pb2.TaskDispatch()
    message.task_id = task_id
    message.student_id = student_id
    message.files_hash = files_hash
    message.data.extend(files_message_content)

    sender_channel.basic_publish(body=files_message.SerializeToString(), routing_key=FILESERVER_ROUTING_KEY, exchange=EXCHANGE_NAME)
    sender_channel.basic_publish(body=message.SerializeToString(), routing_key=QUEUE_TO_WORKER_NAME, exchange=EXCHANGE_NAME)
    time.sleep(0.01)


app = Flask(__name__)


@app.get("/evaluator")
def send_evaluator_message():
    print("Received call")
    amount = int(request.args.get('amount'))
    for _ in range(amount):
        send_message("1", "1", b"1", [".gitignore"])

    return "Message sent"


if __name__ == '__main__':
    setup_channel()
    thread = threading.Thread(target=start_consuming)
    thread.start()
    time.sleep(1)
    print("Evaluator producer - start")

    app.run(host="0.0.0.0", port=5000)
