FROM golang:1.19-alpine

RUN apk update && apk add --no-cache git curl

RUN mkdir /app

COPY . /app

RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o /bin/app.

RUN CGO_ENABLED=0

WORKDIR /app


RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin


CMD ["air"]