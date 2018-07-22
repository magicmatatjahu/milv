FROM golang:1.10 as builder
WORKDIR /milv
ENV SRC_DIR=/go/src/github.com/magicmatatjahu/milv/
ADD . $SRC_DIR
RUN go get gopkg.in/yaml.v2
RUN go get github.com/olekukonko/tablewriter
RUN go get github.com/pkg/errors
RUN cd $SRC_DIR; CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .; cp main /milv/

FROM golang:alpine
LABEL Maintainer Maciej Urbańczyk <github.com/magicmatatjahu>
LABEL source=git@github.com:magicmatatjahu/milv.git
RUN apk update && apk add bash
WORKDIR /milv
COPY --from=builder /milv/main /milv/
ENTRYPOINT ["./main"]