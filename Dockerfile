FROM public.ecr.aws/docker/library/golang:1.20.4
USER root
WORKDIR /lapson_go_api_sample
COPY . .
RUN go mod tidy
RUN go build -o backend ./cmd/game/main.go
# COPY ./config_prod.json ./config.json
EXPOSE 3000
CMD ["/lapson_go_api_sample/backend"]