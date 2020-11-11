package org.news.newstrackerapi.controller;

import org.news.newstrackerapi.model.Article;
import org.news.newstrackerapi.repository.NewsRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/news")
public class NewsTrackerController {

    @Autowired
    private NewsRepository repository;

    @RequestMapping(value = "/getAll", method = RequestMethod.GET)
    public List<Article> getAllArticles() {
        return repository.findAll();
    }
}
