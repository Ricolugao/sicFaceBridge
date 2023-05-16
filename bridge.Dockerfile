FROM golang:1.20.1

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

COPY go.mod go.sum ./
RUN go mod download && go mod verify

RUN apt-get update && \
    apt-get install -y build-essential wget git && \
    wget https://github.com/edenhill/librdkafka/archive/refs/tags/v1.8.0.tar.gz && \
    tar xzf v1.8.0.tar.gz && \
    cd librdkafka-1.8.0 && \
    ./configure --prefix /usr && \
    make && \
    make install && \
    cd .. && \
    rm -rf librdkafka-1.8.0 && \
    rm v1.8.0.tar.gz && \
    ldconfig
    

CMD ["tail", "-f", "/dev/null"]