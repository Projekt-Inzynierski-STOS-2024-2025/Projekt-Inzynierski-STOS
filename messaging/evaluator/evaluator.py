from pika import BlockingConnection
from pika import ConnectionParameters
import random
import threading
import time
import messages_pb2


HOST_NAME: str = 'localhost'
EXCHANGE_NAME: str = 'stos'
QUEUE_TO_WORKER_NAME: str = 'ev_tasks'
QUEUE_FROM_WORKER_NAME: str = 'ev_files'


connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
channel = connection.channel()


def setup_channel():
    channel.exchange_declare(exchange=EXCHANGE_NAME, exchange_type='direct')
    channel.queue_declare(queue=QUEUE_TO_WORKER_NAME)
    channel.queue_declare(queue=QUEUE_FROM_WORKER_NAME)

    channel.queue_bind(exchange=EXCHANGE_NAME, queue=QUEUE_TO_WORKER_NAME, routing_key=QUEUE_TO_WORKER_NAME)
    channel.queue_bind(exchange=EXCHANGE_NAME, queue=QUEUE_FROM_WORKER_NAME, routing_key=QUEUE_FROM_WORKER_NAME)


def resolve_message(ch, parameters, method, body):
    message = messages_pb2.TaskDispatch()
    message.ParseFromString(body)
    time.sleep(random.random() * 2)

    print("---Message---")
    print(message.student_id)
    print(message.task_id)
    print(message.files_hash)
    print(message.data)
    print("---Message end---")


def start_consuming():
    channel.basic_consume(queue=QUEUE_FROM_WORKER_NAME, on_message_callback=resolve_message)
    print("Evaluator consumer - start")
    channel.start_consuming()


def send_message(task_id: str, student_id: str, files_hash: bytes, files: list):
    files_message = []
    for filename in files:
        file_message = messages_pb2.File()
        file_message.name = filename
        file_message.data = open(filename, "rb").read()
        files_message.append(file_message)

    message = messages_pb2.TaskDispatch()
    message.task_id = task_id
    message.student_id = student_id
    message.files_hash = files_hash
    message.data.extend(files_message)

    channel.basic_publish(body=message.SerializeToString(), routing_key=QUEUE_TO_WORKER_NAME, exchange=EXCHANGE_NAME)


if __name__ == '__main__':
    thread = threading.Thread(target=start_consuming)
    thread.start()
    time.sleep(1)
    print("Evaluator producer - start")

    while True:
        pub = input("Press any key to send message")
        send_message("1", "1", b"1", [".gitignore"])
        start_time = time.time()
        print(f"Sent message in {time.time() - start_time} seconds")
