package org.example;

import com.sun.net.httpserver.HttpExchange;
import lombok.SneakyThrows;

public class FixedIntervals implements TrafficManager {
    private int queryAmount;
    private int queryLimit;
    private long intervalTime;
    private boolean isRunning;
    public FixedIntervals(int queryLimit, int intervalTime){
        this.queryLimit = queryLimit;
        this.intervalTime = intervalTime;
    }

    public boolean addRequest(HttpExchange exchange){
        if (queryAmount < queryLimit){
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
    void clear(){
        queryAmount = 0;
    }
    class FixedIntervalsAction implements Runnable{
        @SneakyThrows
        public void run(){
            while(isRunning){
                Thread.sleep(intervalTime * 1000);
                System.out.println("clean");
                clear();
            }
        }
    }
}