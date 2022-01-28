# golang-echo-sample

Make an out-of-the-box backend based on golang-echo

```bash
Echo + MySQL + go-resty + Gorm + ZAP +configor
```

### Tree view

```bash
.
├── application.yml
├── client
│  ├── bible.go
│  ├── bible_test.go
│  ├── models.go
│  └── resp.go
├── docker
│  ├── build.sh
│  └── run.sh
├── Dockerfile
├── docs
│  ├── docs.go
│  ├── swagger.json
│  └── swagger.yaml
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
└── utils
   ├── config
   │  └── config.go
   ├── logger
   │  └── logger.go
   ├── middlewares
   │  ├── logMiddleware.go
   │  └── wrapper.go
   └── util.go
```

## Docker

### how to build

```bash
sh docker/build.sh
```

### how to run

```bash
sh docker/run.sh
```

## References

- [https://gorm.io/](https://gorm.io/)
- [https://echo.labstack.com/](https://echo.labstack.com/)
- [https://github.com/uber-go/zap](https://github.com/uber-go/zap)
- [https://github.com/jinzhu/configor](https://github.com/jinzhu/configor)
- [https://github.com/go-resty/resty](https://github.com/go-resty/resty)
- [https://github.com/ybkuroki/go-webapp-sample](https://github.com/ybkuroki/go-webapp-sample)
- [https://github.com/brpaz/echozap](https://github.com/brpaz/echozap)
- [https://github.com/sandipb/zap-examples](https://github.com/sandipb/zap-examples)


