FROM golang:1.19-alpine

ENV LANGUAGE="en"

ENV APP_HOME /wikipediaTest/src/mathapp
RUN mkdir -p "$APP_HOME"
RUN apk add --no-cache ca-certificates &&\
    chmod +x "$APP_HOME"
WORKDIR "$APP_HOME"
EXPOSE 8010
CMD ./${APP_NAME}