FROM golang:1.15

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o reviewsvc ./cmd/review/main.go ./cmd/review/repository.go

WORKDIR /bin

RUN cp /build/reviewsvc /bin

EXPOSE 8080

CMD [ "/bin/reviewsvc" ]