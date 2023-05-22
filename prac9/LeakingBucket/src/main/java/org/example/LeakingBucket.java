package org.example;

import com.sun.net.httpserver.HttpExchange;
import lombok.SneakyThrows;

import java.util.ArrayDeque;
import java.util.Queue;

public class LeakingBucket implements TrafficManager {
    private Queue<HttpExchange> requestQueue = new ArrayDeque<>(); // очередь HTTP запросов.
    private boolean isRunning = false;
    private int queueCapacity;
    private int handleTime;

    LeakingBucket(int queueCapacity, int handleTime) { // КОНСТРУКТОР ИНИЦИАЛИЗИРУЕТ ПОЛЯ
        this.queueCapacity = queueCapacity;
        this.handleTime = handleTime;
    }
    public boolean addRequest(HttpExchange exchange) {
        if (queueCapacity > requestQueue.size()) {  // проверка на превышение лимита запросов
            requestQueue.add(exchange);
            return true;
        }
        return false;
    }
    public void start() {
        new Thread(new FixedIntervalsAction()).start(); // запуск нового потока
        isRunning = true;
    }
    public void stop() {
        isRunning = false;
    }
    class FixedIntervalsAction implements Runnable{ // вложеный класс
        @SneakyThrows
        @Override
        public void run() { // метод вызываемый в новом потоке в строке 27
            while (isRunning) {
                Thread.sleep(handleTime * 1000); // ждем указанный промежуток времени
                requestQueue.poll(); // освобождаем место для одного нового запроса
                System.out.println("poll");
            }
        }
    }
}