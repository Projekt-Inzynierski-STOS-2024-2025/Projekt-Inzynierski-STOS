package stos.worker;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.ExitCodeGenerator;
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


@SpringBootApplication
@EnableScheduling
@Component
public class WorkerApplication {

	private static final Logger logger = LoggerFactory.getLogger(WorkerApplication.class);
	private final String consumeUrl;
	private final String completeUrl;
	private final Integer taskCompletionTime;
	private final RestTemplate restTemplate;


	public WorkerApplication(
					 @Value("${evaluator.consume.url}") String consumeUrl,
					 @Value("${evaluator.complete.url}") String completeUrl,
					 @Value("${task.completion.time}") Integer taskCompletionTime) {
		this.restTemplate = new RestTemplate();
		this.consumeUrl = consumeUrl;
		this.completeUrl = completeUrl;;
		this.taskCompletionTime = taskCompletionTime;
		if(this.completeUrl == null || this.consumeUrl == null || this.taskCompletionTime == null){
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
            timeConsumingTask(task);
            sendCompletionRequest(task);
        } catch (HttpClientErrorException ex) {
			logger.error("Get request response - {}", ex.getStatusCode().value());
		} catch (Exception ex) {
			logger.error("{}", ex.getMessage());
		}
	}

	public void timeConsumingTask(Task task){
		logger.info("Start processing task {}", task.getId());
		try {
			Thread.sleep(taskCompletionTime);
		} catch (InterruptedException e) {
			logger.error("{}", e.getMessage());
		}
		logger.info("Finished processing task {}", task.getId());
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
