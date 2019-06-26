FROM golang:1.11-alpine as builder
RUN apk --no-cache add git
COPY . /src
WORKDIR /src
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

FROM alpine:latest
RUN apk --no-cache add ca-certificates libcap
COPY --from=builder /app ./
RUN setcap 'cap_net_bind_service=+ep' ./app
USER nobody
EXPOSE 80
ENTRYPOINT ["./app"]
