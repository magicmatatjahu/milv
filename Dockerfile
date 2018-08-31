FROM golang:1.10.3-alpine3.8 as builder

ENV BASE_APP_DIR /go/src/github.com/magicmatatjahu/milv
WORKDIR ${BASE_APP_DIR}

COPY ./ ${BASE_APP_DIR}/

RUN go build -v -o main .
RUN mkdir /app && mv ./main /app/main

FROM alpine:3.8
LABEL Maintainer Maciej Urbańczyk <github.com/magicmatatjahu>
LABEL source = git@github.com:magicmatatjahu/milv.git

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && apk add bash

COPY --from=builder /app /app

ENTRYPOINT ["/app/main"]