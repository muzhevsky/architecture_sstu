package org.example;

import com.sun.net.httpserver.HttpExchange;
import lombok.SneakyThrows;

import java.util.ArrayList;
import java.util.List;

public class FixedIntervals implements TrafficManager {
    private List<HttpExchange> queries;
    private int queryAmount;
    private int queryLimit;
    private long intervalTime;
    private boolean isRunning;
    public FixedIntervals(int queryLimit, int intervalTime){
        this.queryLimit = queryLimit;
        this.intervalTime = intervalTime;
        this.queries = new ArrayList<>(queryLimit);
    }

    public boolean addRequest(HttpExchange exchange){
        if (queryAmount < queryLimit){
            queries.add(exchange);
            queryAmount++;
            return true;
        }
        return false;
    }
    public void start(){
        isRunning = true;
        new Thread(new FixedIntervalsAction()).start();
    }
    public void stop(){
        isRunning = false;
    }
    void clean(){
        queryAmount = 0;
    }
    class FixedIntervalsAction implements Runnable{
        @SneakyThrows
        public void run(){
            while(isRunning){
                Thread.sleep(intervalTime * 1000);
                System.out.println("clear");
                clean();
            }
        }
    }
}