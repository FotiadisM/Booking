FROM golang:1.15

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o listingsvc ./cmd/listing/main.go ./cmd/listing/repository.go

WORKDIR /bin

RUN cp /build/listingsvc /bin

EXPOSE 8080

CMD [ "/bin/listingsvc" ]