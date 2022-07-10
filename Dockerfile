From golang:1.18.3

WORKDIR /echoserver

COPY go.mod ./
COPY go.sum ./
copy pb ./
COPY *.go ./


RUN go mod download
RUN go build -o /frontend

EXPOSE 8080

CMD [ "/frontend" ]