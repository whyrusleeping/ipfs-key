FROM golang:alpine

RUN apk update && \
    apk add git --no-cache && \
    go get github.com/whyrusleeping/ipfs-key

CMD ["ipfs-key"]