package org.news.newstrackerapi.model;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

@Document(collection = "article")
public class Article {


    @Id
    private String id;
    private String author;
    private String title;
    private String description;
    private String url;
    private String urltoimage;
    private String publishedat;
    private String content;
    private Source source;

    public Article() {
    }

    public Article(String author,
                   String title,
                   String description,
                   String url,
                   String urltoimage,
                   String publishedAt,
                   String content,
                   Source source) {
        this.author = author;
        this.title = title;
        this.description = description;
        this.url = url;
        this.urltoimage = urltoimage;
        this.publishedat = publishedat;
        this.content = content;
        this.source = source;
    }

    public String getAuthor() {
        return author;
    }

    public void setAuthor(String author) {
        this.author = author;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getUrl() {
        return url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public String getUrltoimage() {
        return urltoimage;
    }

    public void setUrltoimage(String urltoimage) {
        this.urltoimage = urltoimage;
    }

    public String getPublishedat() {
        return publishedat;
    }

    public void setPublishedat(String publishedat) {
        this.publishedat = publishedat;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Source getSource() {
        return source;
    }

    public void setSource(Source source) {
        this.source = source;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }
}

