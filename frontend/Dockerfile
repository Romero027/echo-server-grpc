From golang:1.18.3

WORKDIR /echoserver

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY frontend.go ./

RUN go build -o /frontend

EXPOSE 8080

CMD [ "/frontend" ]