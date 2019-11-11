FROM golang:alpine as builder
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN apk add --no-cache git && \
    go get -u github.com/gorilla/mux
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

# Run the binary.
ENTRYPOINT ["/app/main"]