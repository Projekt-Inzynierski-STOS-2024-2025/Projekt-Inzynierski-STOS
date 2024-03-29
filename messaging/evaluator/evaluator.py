from pika import BlockingConnection
from pika import ConnectionParameters
import random
import threading
import time
import messages
import messages_pb2


def evaluator_task_callback(ch, method, properties, body):
    print("evaluator_task_callback")
    message = messages_pb2.File()
    message.ParseFromString(body)
    print(message.name)
    print(message.data)
    simulate_workload(body)


def worker_task_callback(ch, method, properties, body):
    print("worker_task_callback")
    simulate_workload(body)


def evaluator_file_callback(ch, method, properties, body):
    print("evaluator_file_callback")
    simulate_workload(body)


def file_evaluator_callback(ch, method, properties, body):
    print("file_evaluator_callback")
    simulate_workload(body)


def simulate_workload(body):
    time_to_wait = random.random() * 3
    time.sleep(time_to_wait)


HOST_NAME: str = 'localhost'
queue_mapping: dict = {
    'evaluator_task_queue': {'exchange': 'evaluator_worker', 'callback': evaluator_task_callback},
    'worker_task_queue': {'exchange': 'evaluator_worker', 'callback': worker_task_callback},
    'evaluator_file_queue': {'exchange': 'evaluator_files', 'callback': evaluator_file_callback},
    'file_evaluator_queue': {'exchange': 'evaluator_files', 'callback': file_evaluator_callback}
}
# TODO: change to correct queue
message_to_queue_mapping: dict = {
    'file_message': 'evaluator_task_queue',
    'files_message': 'evaluator_task_queue',
    'files_info_message': 'evaluator_task_queue',
    'task_dispatch_message': 'evaluator_task_queue',
    'error_message': 'evaluator_task_queue'
}

connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
channel = connection.channel()

# Create exchanges
exchanges = set([config['exchange'] for config in queue_mapping.values()])
for config in exchanges:
    channel.exchange_declare(exchange=config, exchange_type='direct')
    print(f"Created exchange: {config}")

# Create and bind queues
for queue, config in queue_mapping.items():
    channel.queue_declare(queue=queue)
    channel.queue_bind(exchange=config['exchange'], queue=queue, routing_key=queue)
    print(f"Bound {queue} to {config['exchange']}")


def start_consuming():
    for queue, config in queue_mapping.items():
        print(f"Queue {queue} invokes callback: {config['callback']}")
        channel.basic_consume(queue=queue, on_message_callback=config['callback'], auto_ack=True)
    print("Evaluator consumer - start")
    channel.start_consuming()


def send_message(message_callback, data: dict):
    message = getattr(messages, message_callback)(data)
    routing_key = message_to_queue_mapping[message_callback]
    exchange_name = queue_mapping[routing_key]['exchange']

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
    time.sleep(1)
    print("Evaluator producer - start")

    while True:
        message = input("Enter message callback name: ")
        payload = input("Pass message values: ")
        resolved_payload = resolve_payload(payload)

        start_time = time.time()
        send_message(message, resolved_payload)
        print(f"Sent message in {time.time() - start_time} seconds")
