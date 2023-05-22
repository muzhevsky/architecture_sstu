package org.example;

import com.sun.net.httpserver.HttpServer;

import java.net.InetSocketAddress;

public class Main {
    public static void main(String[] args) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0); // запуск сервера по адресу http://localhost:8000/

        var trafficManager = new FixedIntervals(3, 5);
        var intervalHandler = new MyHttpHandler(trafficManager);
        trafficManager.start();

        server.createContext("/", intervalHandler); // для сервера
        server.setExecutor(null);
        server.start();
    }
}
