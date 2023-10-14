### Builder ###
FROM golang:alpine as builder
WORKDIR /go/src/mekari
COPY . .

RUN go mod download
RUN go build -o main .

### Main App ###
FROM alpine:3
COPY --from=builder /go/src/mekari/main .
COPY --from=builder /go/src/mekari/.env .

EXPOSE 8081
CMD ["./main"]