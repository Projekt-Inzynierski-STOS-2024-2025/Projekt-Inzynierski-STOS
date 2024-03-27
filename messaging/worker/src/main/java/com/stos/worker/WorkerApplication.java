package com.stos.worker;
import java.util.Scanner;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class WorkerApplication implements CommandLineRunner {

	@Autowired
	private Worker worker;

	@Autowired
	EvaluatorMock evaluatorMock;

	private final Scanner scanner;

	public WorkerApplication() {
		this.scanner = new Scanner(System.in);
	}

	public static void main(String[] args) {
		SpringApplication.run(WorkerApplication.class, args);
	}

	@Override
	public void run(String... args) throws Exception {
		while (true) {
			String input = scanner.nextLine();
			if ("q".equalsIgnoreCase(input)) {
				break;
			} else if ("wp".equalsIgnoreCase(input)) {
				worker.sendGeneratedMessageRabbit();
			} else if ("wc".equalsIgnoreCase(input)) {
				System.out.println(worker.receiveMessageRabbit());
			} else if ("ep".equalsIgnoreCase(input)) {
				evaluatorMock.sendGeneratedMessageRabbit();
			} else if ("ec".equalsIgnoreCase(input)) {
				System.out.println(evaluatorMock.receiveMessageRabbit());
			} else if ("wpk".equalsIgnoreCase(input)) {
				worker.sendGeneratedMessageKafka();
			} else if ("eck".equalsIgnoreCase(input)) {
				evaluatorMock.readFromKafkaTopic();
			}
		}
	}
}