package org.news.newstrackerapi;

import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.data.mongodb.core.MongoTemplate;

@Configuration
public class MongoTemplateConfiguration {

    @Value("${MONGODB-HOST}")
    private String mongoDbHost;
    @Value("${MONGODB-NAME}")
    private String mongoDbName;
    @Value("${MONGODB-PORT}")
    private String mongoDbPort;

    public @Bean
    MongoTemplate mongoTemplate() {
        return new MongoTemplate(mongoClient(), mongoDbName);
    }

    public @Bean
    MongoClient mongoClient() {
        return MongoClients.create(String.format("mongodb://%s:%s", mongoDbHost, mongoDbPort));
    }
}
