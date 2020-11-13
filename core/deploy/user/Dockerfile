FROM golang:1.15

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o usersvc ./cmd/user/main.go ./cmd/user/repository.go

WORKDIR /bin

RUN cp /build/usersvc /bin

EXPOSE 8080

CMD [ "/bin/usersvc" ]