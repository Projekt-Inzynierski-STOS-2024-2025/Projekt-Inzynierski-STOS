package stos.worker;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;
import org.springframework.web.client.HttpClientErrorException;
import org.springframework.web.client.RestTemplate;
import java.io.IOException;


@SpringBootApplication
@EnableScheduling
@Component
public class WorkerApplication {

	private static final Logger logger = LoggerFactory.getLogger(WorkerApplication.class);
	private final String consumeUrl;
	private final String completeUrl;
	private final RestTemplate restTemplate;


	public WorkerApplication(
					 @Value("${evaluator.consume.url}") String consumeUrl,
					 @Value("${evaluator.complete.url}") String completeUrl) {
		this.restTemplate = new RestTemplate();
		this.consumeUrl = consumeUrl;
		this.completeUrl = completeUrl;

		if(this.completeUrl == null || this.consumeUrl == null ){
			logger.error("Not all env set");
		}
	}

	public static void main(String[] args) {
		SpringApplication.run(WorkerApplication.class, args);
		logger.info("Worker: Application started");
	}


	@Scheduled(fixedDelay = 3000)
	public void sendGetRequest(){
		logger.info("Sending get request at: " + consumeUrl);
		try {
			Task task = new Task(restTemplate.getForObject(consumeUrl, String.class));
            logger.info("Received task with UUID - {}", task.getId());
            String time = timeConsumingTask(task);
			task.setTime(time);
            sendCompletionRequest(task);
        } catch (HttpClientErrorException ex) {
			logger.error("Get request response - {}", ex.getStatusCode().value());
		} catch (Exception ex) {
			logger.error("{}", ex.getMessage());
		}
	}

	public String timeConsumingTask(Task task){
		long startTime = System.currentTimeMillis();
		logger.info("Start processing task {}", task.getId());
		try {
			Process process = Runtime.getRuntime().exec("workload");
			process.waitFor();
		} catch (IOException e) {
			logger.error("{}", e.getMessage());
		} catch (InterruptedException e) {
			e.printStackTrace();
		}
		logger.info("Finished processing task {}", task.getId());
		long endTime = System.currentTimeMillis();
		return String.valueOf(endTime - startTime);
	}

	public void sendCompletionRequest(Task task){
		HttpHeaders headers = new HttpHeaders();
		headers.setContentType(MediaType.APPLICATION_JSON);
		ObjectMapper objectMapper = new ObjectMapper();
		String jsonBody = null;
		try {
			jsonBody = objectMapper.writeValueAsString(task);
		} catch (Exception e) {
			logger.error("Error converting TaskData to JSON - {}", e.getMessage());
		}
		HttpEntity<String> request = new HttpEntity<>(jsonBody, headers);
		logger.info("Sending post request at {}", completeUrl);
		restTemplate.postForObject(completeUrl, request, Void.class);
	}
}
