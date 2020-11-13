FROM golang:1.15

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o searchconsumersvc ./cmd/search_consumer/main.go ./cmd/search_consumer/repository.go

WORKDIR /bin

RUN cp /build/searchconsumersvc /bin

EXPOSE 8080

CMD [ "/bin/searchconsumersvc" ]