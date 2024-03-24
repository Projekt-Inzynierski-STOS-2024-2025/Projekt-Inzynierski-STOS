docker compose up -d

timeout /t 10
start python.exe evaluator-consumer.py

timeout /t 5
python.exe evaluator-producer.py 30
timeout /t 5