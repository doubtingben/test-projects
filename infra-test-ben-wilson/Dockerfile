FROM golang:1.9

WORKDIR /go/src/github.com/INFURA/infra-test-ben-wilson
COPY . .

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080

ENTRYPOINT ["infra-test-ben-wilson"]
