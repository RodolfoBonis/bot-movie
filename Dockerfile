FROM golang:1.19.2 AS build-env

ENV CGO_ENABLED 0

RUN apt-get install git

WORKDIR /go/src/github.com/RodolfoBonis/bot_movie/
ADD . /go/src/github.com/RodolfoBonis/bot_movie/

RUN go build -o bot_movie -v /go/src/github.com/RodolfoBonis/bot_movie/

COPY . ./

FROM alpine:3.15

WORKDIR /app

COPY --from=build-env /go/src/github.com/RodolfoBonis/bot_movie /app/

CMD ["/app/bot_movie"]