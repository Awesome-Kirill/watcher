## Task
<img align="right" width="50%" src="./images/big-gopher.jpg">
Write a program, that will check the list of domains for accessibility

Once a minute we need to check, if domains are available and get their ping.
We have many users, who want to know domains ping.
Users have three options:

1. Get the time for specific domain
2. Get the domain with min ping
3. Get the domain with max ping

Both we have administrators, who want to get statistics for this endpoints


## HOWTO
- :running_man: build
```
docker build . --tag=watcher:1.0
```
- :running_man: run
```
docker run -it -p 8080:8080 --rm --name app-watcher watcher:1.0
```

```
curl --location --request GET 'http://localhost:8080/stat/reddit.com/site'
```

```
curl --location --request GET 'http://localhost:8080/stat/min'
```

```
curl --location --request GET 'http://localhost:8080/stat/max'
```

```
curl --location --request GET 'http://localhost:8080/admin/stat' \
  --header 'Authorization: Bearer any-secret'
```
- :test_tube: run tests with `make test`
- :sunflower: run linter with `make lint`
- :heart: generate documentation with `make swagger`
- see docs in http://localhost:8080/swagger/
