package org.news.newstrackerapi.repository;


import org.news.newstrackerapi.model.Article;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.List;

public interface NewsRepository extends MongoRepository<Article, String> {

    List<Article> findByAuthor(String author);
}
