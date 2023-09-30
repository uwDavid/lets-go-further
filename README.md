# Let's Go Further - Movie Database API
Follow along of the book `Let's Go Further`
We are building a Movie database API. 

## Running This App
#### Generate TLS Self-signed Certificate
```
mkdir tls
cd tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

#### Running MySQL container
```
docker compose up -d
```
However, you may get this error when trying to connect using `user@localhost`. 
There are 2 ways to solve this: 
1. Add `MYSQL_ROOT_HOST: '%'` environment variable to allow login from any IP. 
Caveat is that you have to clear the docker volume

2. Use `docker exec` and change the grant table: 
```
SELECT host, user from user;
CREATE USER 'test_web'@'%' IDENTIFIED BY 'pass';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'%';
```