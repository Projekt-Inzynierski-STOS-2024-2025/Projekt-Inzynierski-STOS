FROM maven:3.9.6-eclipse-temurin-21 AS builder

WORKDIR /app

COPY pom.xml .
COPY src src

RUN mvn clean package

FROM rust:1.67 as rust_builder
WORKDIR /usr/src/app
COPY workload/src src
COPY workload/Cargo.toml .
COPY workload/Cargo.lock .
RUN cargo build --locked --release

FROM eclipse-temurin:21

WORKDIR /app
COPY --from=builder /app/target/worker-0.0.1-SNAPSHOT.jar /app/worker.jar
COPY --from=rust_builder /usr/src/app/target/release/workload /usr/local/bin/workload


ENV SPRING_APPLICATION_NAME=worker
ENV LOGGING_LEVEL_ORG_SPRINGFRAMEWORK=WARN
ENV SPRING_MAIN_LOG-STARTUP-INFO=false
ENV SPRING_MAIN_BANNER-MODE=OFF
ENV EVALUATOR_CONSUME_URL=http://evaluator:2137/consume
ENV EVALUATOR_COMPLETE_URL=http://evaluator:2137/complete
ENV LOG_HOST=logstash-logstash
ENV LOG_PORT=5228

CMD ["java", "-jar", "worker.jar"]
