FROM golang:1.15

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bookingsvc ./cmd/booking/main.go ./cmd/booking/repository.go

WORKDIR /bin

RUN cp /build/bookingsvc /bin

EXPOSE 8080

CMD [ "/bin/bookingsvc" ]