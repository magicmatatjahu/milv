FROM golang:alpine

LABEL Maintainer Maciej Urbańczyk <github.com/magicmatatjahu>
LABEL source=git@github.com:magicmatatjahu/milv.git

RUN apk update && apk add bash && apk add git
RUN go get -u -v github.com/magicmatatjahu/milv

VOLUME /milv
WORKDIR /milv

ENTRYPOINT ["milv"]