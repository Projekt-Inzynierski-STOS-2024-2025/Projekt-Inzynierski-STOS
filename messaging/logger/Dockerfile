FROM golang:1.18
COPY . /build
WORKDIR /build
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /logger
CMD /logger
