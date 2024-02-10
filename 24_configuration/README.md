#  Configuration

## Add Configuration for config go via env
```go

    err := godotenv.Load("default.env")
    if err !=nill {
        log.Println("Please use the configuration file / environment variables: %s", err)
    }
    ...    

```

## How to use configuration via environment variable / environment file
```go
    os.Getenv("<ENV_KEY>")
```

