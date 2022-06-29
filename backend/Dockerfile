FROM golang:1.17.6

RUN go install github.com/beego/bee/v2@latest
ENV GO111MODULE=on

ENV APP_HOME /go/src/starfly
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
COPY . "$APP_HOME"

EXPOSE 9999

CMD ["bee", "run"]