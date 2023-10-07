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

#### Running Test SMTP Server
Instead of using `Mailtrap`, we will just run a `MailHog` container. 
```
docker run -d -p 1025:1025 -p 8025:8025 mailhog/mailhog
```
We change the config to point the host to `localhost` and port to `1025`

#### Running PSQL container
```
docker compose up -d
```
