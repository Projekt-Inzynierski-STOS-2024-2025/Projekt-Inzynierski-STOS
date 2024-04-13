package com.stos.worker;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class WorkerApplication implements CommandLineRunner {

	@Autowired
	private Worker worker;
	public static void main(String[] args) {
		SpringApplication.run(WorkerApplication.class, args);
	}
	
	@Override
	public void run(String... args) {
		try{
			worker.workerTest();
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}