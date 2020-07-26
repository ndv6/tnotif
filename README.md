# Documentation

## API Contract

### [POST] /sendMail
Sending mail and log it into database

Request :
```
{
    "email" : "testing@example.com",
    "token" : "SecretToken"
}
```

Response <br/>
200 Ok
```
{
    "status" : "SUCCESS",
    "message" : "Send Mail Success",
    "data" : {
        "email" : "testing@example.com"
    }
}
```
400 Bad Request
```
{
    "error" : "the function is returning error"
}
```

## How to run
1. Set your config in `config.json` file
2. ```go run main.go```

## How to run unit test
1. `<env> go test ./...` 