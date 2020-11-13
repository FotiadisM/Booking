FROM golang:1.15 as Builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o reviewsvc ./cmd/review/main.go ./cmd/review/repository.go

FROM alpine:3.12.1

WORKDIR /bin
RUN apk --no-cache add ca-certificates
COPY --from=Builder /build/reviewsvc .

EXPOSE 8080

CMD [ "/bin/reviewsvc" ]