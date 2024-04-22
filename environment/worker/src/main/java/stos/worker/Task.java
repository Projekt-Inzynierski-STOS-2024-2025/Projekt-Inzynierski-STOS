package stos.worker;
import com.fasterxml.jackson.annotation.JsonProperty;

public class Task {
    @JsonProperty("id")
    private String id;

    @JsonProperty("time")
    private String time;

    public Task(String id) {
        this.id = id;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public void setTime(String time){ this.time = time;}
}


