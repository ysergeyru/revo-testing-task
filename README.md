# Revo testing service

## Task description

Implement a microservice on Golang, which allows to understand how many threads you can parse a particular site without errors (load testing).

The service receives a search string at the entrance, for example, "buy playstation". From the Yandex search results (implemented parser [here](https://github.com/kkhrychikov/revo-testing/blob/main/serp.go)) we get a list of urls. Further, for each url, you need to conduct a small benchmark - how many parallel requests from one IP this url can handle without errors. The maximum response time is up to 3 seconds. The response to the original GET should be the map "host" => "recommended number of concurrent requests".

**Mandatory points to be implemented**:

- the service must be wrapped in docker
- interaction through one endpoint GET /sites?search=foobar
- with cold cache, the service should respond no longer than 30 seconds.
- setting parameters via config
