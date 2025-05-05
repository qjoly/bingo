FROM golang:1.23.4-bookworm as builder
COPY go.mod /app/
WORKDIR /app
RUN go mod tidy
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /app/bingo
FROM scratch
COPY --from=builder /app/bingo /bingo
ENTRYPOINT ["/bingo"]
