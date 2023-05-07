FROM golang:alpine3.17
RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o fazri1/go-apotek
EXPOSE 8080
CMD ["/fazri1/go-apotek"]
