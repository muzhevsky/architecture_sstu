package org.example;

import com.google.gson.Gson;
import com.mysql.cj.jdbc.MysqlDataSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.jdbc.repository.config.EnableJdbcRepositories;

import javax.sql.DataSource;

@Configuration
@EnableJdbcRepositories("org")
public class Config {

    @Bean("dataSource")
    public DataSource dataSource() {
        var datasource = new MysqlDataSource();

        datasource.setDatabaseName("marketplace");
        datasource.setServerName("localhost");
        datasource.setPort(3306);
        datasource.setUser("root");
        datasource.setPassword("MysqlPassssaP1");

        return datasource;
    }

    @Bean("gson")
    public Gson gson(){
        return new Gson();
    }
}
