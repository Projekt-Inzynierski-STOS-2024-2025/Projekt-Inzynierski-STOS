from pika import BlockingConnection
from pika import ConnectionParameters
import time
import messages


HOST_NAME: str = 'localhost'
queue_exchange_mapping: dict = {
    'evaluator_worker': ['evaluator_task_queue', 'worker_task_queue'],
    'evaluator_files': ['evaluator_file_queue', 'file_evaluator_queue']
}

connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
channel = connection.channel()

for exchange, queues in queue_exchange_mapping.items():
    channel.exchange_declare(exchange=exchange, exchange_type='direct')
    for queue in queues:
        channel.queue_declare(queue=queue)
        channel.queue_bind(exchange=exchange, queue=queue, routing_key=queue)
        print(f"Bound {queue} to {exchange}")


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
