FROM golang:1.20.1

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y 
    

CMD ["tail", "-f", "/dev/null"]