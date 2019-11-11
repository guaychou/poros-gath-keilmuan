FROM golang:alpine as builder
RUN apk add --no-cache git 
WORKDIR /go/src/github.com/go-redis/redis/
RUN git clone https://github.com/go-redis/redis.git .
RUN mkdir /app
ADD . /app/
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main
RUN adduser -S -D -H -h /app appuser


FROM scratch
# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /app/main /app/main

# Use an unprivileged user.
USER appuser
EXPOSE 8080

# Run the hello binary.
ENTRYPOINT ["/app/main"]