
# Logging Example
This repo is created for the [Ultimate Debugging with Go](https://www.bytesizego.com/the-ultimate-guide-to-debugging-with-go) Course that is currently in waitlist.

The goal of this project is to give an example of how you might structure your Go project, configure and use structured logging, send them to Elastic Search and view them by Kibana.  In this instance, I send them to Elastic Search via http endpoint (check out log/multi.go to see how I do this).

## Getting Started with this project
Firstly, you must have docker installed. Once that is done you should be able to run `docker-compose up` in the root of the project to start elastic search and kibana.

Kibana will be available to you on `http://localhost:5601`. Elasticsearch will be available at `http://localhost:9200`.   The first time you visit Kibana after sending a log, it might ask you to create an index. You can do this by just agreeing with the defaults.

I have included an `.env` file that has the two environment variables you will need. Setting `ENV_LOG_LEVEL` to debug will increase the log level in the project, anything else will
cause it to only error log.

The goal of the project is to replicate a library system. The project is not complete and one of the exercises in the course is to complete it.

You can make requests to the server once running by using:
```
curl --location 'http://localhost:8080/books'
```

There is also a second binary, inside seeder/main.go that will send as many requests as it can to the server until you terminate the binary. You can use this to put a bunch of data into Kibana for you to look at.

The service has a mock database that will error 33% of the time, leading to some error and info logs.

Hope you find this project useful!
