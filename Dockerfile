#FROM golang:1.17.8
FROM golang:alpine3.15 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /app .
ENTRYPOINT ["/main"]