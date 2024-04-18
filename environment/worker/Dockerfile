FROM maven:3.9.6-eclipse-temurin-21 AS builder

WORKDIR /app

COPY pom.xml .
COPY src src

RUN mvn clean package

FROM eclipse-temurin:21

WORKDIR /app
COPY --from=builder /app/target/worker-0.0.1-SNAPSHOT.jar /app/worker.jar

ENV SPRING_APPLICATION_NAME=worker
ENV LOGGING_LEVEL_ORG_SPRINGFRAMEWORK=WARN
ENV SPRING_MAIN_LOG-STARTUP-INFO=false
ENV SPRING_MAIN_BANNER-MODE=OFF
ENV EVALUATOR_CONSUME_URL=http://evaluator:2137/consume
ENV EVALUATOR_COMPLETE_URL=http://evaluator:2137/complete
ENV TASK_COMPLETION_TIME=3000

CMD ["java", "-jar", "worker.jar"]