FROM python:3.11

WORKDIR /app

COPY ./evaluator /app

RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "evaluator.py" ]
