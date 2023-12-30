FROM alpine:latest AS builder

WORKDIR /app
COPY . .

RUN /bin/sh -c "apk add --no-cache go && \
                go mod tidy && \
                go build -o main ."

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 80
ENTRYPOINT ["./main"]

# building the docker image on Unix systems: sudo docker build -t go-hello-world-service:0.1.0 .
# run the docker container on Unix systems: sudo docker run --rm -p 5002:80 -d go-hello-world-service:0.1.0 