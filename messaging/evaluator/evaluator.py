from pika import BlockingConnection
from pika import ConnectionParameters
import random
import threading
import time
import messages
import messages_pb2


def evaluator_task_callback(ch, method, properties, body):
    print("evaluator_task_callback")
    simulate_workload()
    message = messages_pb2.File()
    message.ParseFromString(body)
    print(message.name)
    print(message.data)


def worker_task_callback(ch, method, properties, body):
    print("worker_task_callback")
    simulate_workload()
    message = messages_pb2.File()
    message.ParseFromString(body)
    print(message.name)
    print(message.data)


def evaluator_file_callback(ch, method, properties, body):
    print("evaluator_file_callback")
    simulate_workload()
    message = messages_pb2.File()
    message.ParseFromString(body)
    print(message.name)
    print(message.data)


def file_evaluator_callback(ch, method, properties, body):
    print("file_evaluator_callback")
    simulate_workload()
    message = messages_pb2.File()
    message.ParseFromString(body)
    print(message.name)
    print(message.data)


def simulate_workload():
    time_to_wait = random.random() * 3
    time.sleep(time_to_wait)


HOST_NAME: str = 'localhost'
queue_exchange_mapping: dict = {
    'evaluator_worker': ['evaluator_task_queue', 'worker_task_queue'],
    'evaluator_files': ['evaluator_file_queue', 'file_evaluator_queue']
}
queue_callback_mapping: dict = {
    'evaluator_task_queue': evaluator_task_callback,
    'worker_task_queue': worker_task_callback,
    'evaluator_file_queue': evaluator_file_callback,
    'file_evaluator_queue': file_evaluator_callback
}

connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
channel = connection.channel()

for exchange, queues in queue_exchange_mapping.items():
    channel.exchange_declare(exchange=exchange, exchange_type='direct')
    for queue in queues:
        channel.queue_declare(queue=queue)
        channel.queue_bind(exchange=exchange, queue=queue, routing_key=queue)
        print(f"Bound {queue} to {exchange}")


def start_consuming():
    for queue, callback_function in queue_callback_mapping.items():
        print(f"Queue {queue} invokes callback: {callback_function}")
        channel.basic_consume(queue=queue, on_message_callback=callback_function, auto_ack=True)
    print("Evaluator consumer - start")
    channel.start_consuming()


def send_message(exchange_name: str, routing_key: str, message_callback, data: dict, amount: int = 1):
    message = getattr(messages, message_callback)(data)
    for _ in range(amount):
        channel.basic_publish(body=message.SerializeToString(), routing_key=routing_key, exchange=exchange_name)


def resolve_payload(payload: str) -> dict:
    result = {}
    for pair in payload.split(","):
        key_val = pair.split(":")
        result[key_val[0]] = key_val[1]

    return result


if __name__ == '__main__':
    thread = threading.Thread(target=start_consuming)
    thread.start()
    time.sleep(3)
    print("Evaluator producer - start")

    while True:
        exchange_name = input("Enter exchange name: ")
        routing = input("Enter routing name: ")
        message = input("Enter message callback name: ")
        payload = input("Pass message values: ")
        resolved_payload = resolve_payload(payload)

        start_time = time.time()
        send_message(exchange_name, routing, message, resolved_payload)
        print(f"Sent message in {time.time() - start_time} seconds")
