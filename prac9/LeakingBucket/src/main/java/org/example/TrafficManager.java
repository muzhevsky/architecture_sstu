package org.example;

import com.sun.net.httpserver.HttpExchange;

public interface TrafficManager {
    boolean addRequest(HttpExchange exchange);
    void start();
    void stop();
}
