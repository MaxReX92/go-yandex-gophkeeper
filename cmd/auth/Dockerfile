FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin ./cmd/auth/main.go

FROM scratch
COPY --from=builder /app/bin /app
EXPOSE 8080
CMD [ "/app" ]