# Build stage
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main .

CMD ["/app/main"]
#
#
## Final stage
#FROM alpine:3.13
#
#COPY --from=builder /app/main /app/main
#
#RUN chmod +x /app/main
#
#CMD ["/app/main"]
