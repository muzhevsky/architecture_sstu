package org.example;

import com.sun.net.httpserver.HttpExchange;
import lombok.SneakyThrows;

import java.util.ArrayDeque;
import java.util.Queue;

public class LeakingBucket implements TrafficManager {
    private Queue<HttpExchange> requestQueue = new ArrayDeque<>();
    private boolean isRunning = false;
    private int queueCapacity;
    private int handleTime;

    LeakingBucket(int queueCapacity, int handleTime) {
        this.queueCapacity = queueCapacity;
        this.handleTime = handleTime;
    }
    public boolean addRequest(HttpExchange exchange) {
        if (queueCapacity > requestQueue.size()) {
            requestQueue.add(exchange);
            return true;
        }
        return false;
    }
    public void start() {
        new Thread(new FixedIntervalsAction()).start();
        isRunning = true;
    }
    public void stop() {
        isRunning = false;
    }
    class FixedIntervalsAction implements Runnable{
        @SneakyThrows
        @Override
        public void run() {
            while (isRunning) {
                Thread.sleep(handleTime * 1000);
                requestQueue.poll();
                System.out.println("poll");
            }
        }
    }
}