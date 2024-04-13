FROM maven:3.9.6-eclipse-temurin-21 AS builder

WORKDIR /app

COPY pom.xml .
COPY src src

RUN mvn clean package

FROM eclipse-temurin:21

WORKDIR /app
COPY --from=builder /app/target/worker-0.0.1-SNAPSHOT.jar /app/worker.jar

ENV SPRING_APPLICATION_NAME=worker
ENV SPRING_RABBITMQ_HOST=rabbitmq
ENV SPRING_RABBITMQ_PORT=5672
ENV SPRING_RABBITMQ_USERNAME=guest
ENV SPRING_RABBITMQ_PASSWORD=guest

CMD ["java", "-jar", "worker.jar"]