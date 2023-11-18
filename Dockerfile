FROM golang:alpine as build
RUN apk add --no-cache ca-certificates openssh-client

ARG SSH_PRIVATE_KEY

WORKDIR /go/src/github.com/RapidCodeLab/rapid-prebid-server
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-extldflags "-static"' -o server ./cmd/app

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt \
     /etc/ssl/certs/ca-certificates.crt

COPY --from=build /go/src/github.com/RapidCodeLab/rapid-prebid-server/server \ 
    /server

ENTRYPOINT ["/server"]
