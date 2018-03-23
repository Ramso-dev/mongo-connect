FROM golang:1.8




RUN mkdir -p /go/src/Tools/mongodb-connect/
WORKDIR /go/src/Tools/mongodb-connect/

COPY . /go/src/Tools/mongodb-connect/
RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE 8080