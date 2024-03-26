package com.stos.worker;
import java.util.Scanner;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class WorkerApplication implements CommandLineRunner {

	@Autowired
	private RabbitMqWorker rabbitMqWorker;

	@Autowired
	RabbitMqEvaluatorMock rabbitMqEvaluatorMock;

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
				rabbitMqWorker.sendGeneratedMessage();
			} else if ("wc".equalsIgnoreCase(input)) {
				System.out.println(rabbitMqWorker.receiveMessage());
			} else if ("ep".equalsIgnoreCase(input)) {
				rabbitMqEvaluatorMock.sendGeneratedMessage();
			} else if ("ec".equalsIgnoreCase(input)) {
				System.out.println(rabbitMqEvaluatorMock.receiveMessage());
			}
		}
	}
}