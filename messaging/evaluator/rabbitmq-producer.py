from pika import BlockingConnection
from pika import ConnectionParameters
import time
import messages


def send_message(exchange_name: str, queue_name: str, message_callback, data: dict, amount: int = 1):
    HOST_NAME = 'localhost'
    connection = BlockingConnection(ConnectionParameters(host=HOST_NAME))
    channel = connection.channel()
    channel.exchange_declare(exchange=exchange_name, exchange_type='direct')
    channel.queue_declare(queue=queue_name)
    channel.queue_bind(exchange=exchange_name, queue=queue_name, routing_key='')

    message = getattr(messages, message_callback)(data)
    for _ in range(amount):
        channel.basic_publish(body=message.SerializeToString(), routing_key='', exchange=exchange_name)

    connection.close()


def resolve_payload(payload: str) -> dict:
    result = {}

    for pair in payload.split(","):
        key_val = pair.split(":")
        result[key_val[0]] = key_val[1]

    return result


if __name__ == '__main__':
    print("Evaluator producer - start")
    exchange_name = "test-exchange"#sys.argv[1]
    queue_name = "test-queue"#sys.argv[2]

    while True:
        message = input("Enter message callback name: ")
        payload = input("Pass message values")
        resolved_payload = resolve_payload(payload)

        start_time = time.time()
        send_message(exchange_name, queue_name, message, resolved_payload)
        print(f"Sent message in {time.time() - start_time} seconds")
