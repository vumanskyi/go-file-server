FROM golang:1.15

MAINTAINER Vladyslav Umanskyi <vladumanskyi@gmail.com>

WORKDIR /server

ADD . .

RUN go build -o /bin/server

FROM gcr.io/distroless/base

COPY --from=0 /bin/server /bin/server

ENTRYPOINT [ "/bin/server" ]