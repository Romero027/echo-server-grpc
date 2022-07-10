From golang:1.18.3

WORKDIR /echoserver

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY frontend.go ./

RUN go build -o /server

EXPOSE 9000

CMD [ "/server" ]