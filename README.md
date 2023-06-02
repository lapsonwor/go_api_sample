# go api backend

## Setup

```bash
# link src go path
ln -s $(pwd) /Users/$USER/go/src/lapson_go_api_sample
# Install go package
go mod tidy
# Start project
API_ENV=DEBUG GIN_MODE=debug go run cmd/main.go --config config/config.json
```

## Program Structure

```go
backend
└── cmd/game/main.go /* Responsible for command line options. */
    └── core/core.go /* Responsible for importing all required packages based on configuration. */
        │── controller/controller.go /* Responsible for controlling data management and integration.*/
        │── api/handler/handler.go /* Responsible for handling inbound management and forwarding to the controller. */
        │── api/api.go /* Responsible for providing port monitoring and forwarding services. */
└── log/logger.go /* Responsible for creating and writing log files. */
└── config/config.go /* Responsible for reading and decoding configuration files. */
└── model /* Responsible for saving and decoding data structure and converting it to our database structure. */
```
