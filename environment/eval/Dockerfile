FROM golang:1.22
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /evaluator
ENV LOGSTASH_ADDR=logstash
CMD /evaluator
