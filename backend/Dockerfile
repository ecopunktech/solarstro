# Build stage
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/main

# Final stage
FROM scratch

# copy certificates
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=build /app/main /app/main

EXPOSE 8080

ENTRYPOINT ["/app/main"]