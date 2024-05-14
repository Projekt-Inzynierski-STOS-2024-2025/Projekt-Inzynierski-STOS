import requests
import json

logstash_url = 'http://localhost:9601'  # Update with your Logstash IP and port


def send_log(message):
    headers = {'Content-Type': 'application/json'}
    payload = {'message': message}
    response = requests.post(logstash_url, data=json.dumps(payload))
    if response.status_code == 200:
        print("Log sent successfully")
    else:
        print("Failed to send log")


# Example usage:
send_log("This is a test log message from Python")
