FROM golang:1.10.5 as builder
WORKDIR /GoProject/src/dreba
ENV GOPATH=/GoProject

COPY ./ /GoProject/src/dreba/
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o  _output/dreba ./dreba.go

FROM alpine:latest as prod
LABEL maintainers="longhui.li"
LABEL description="dreba"
ARG version
ENV COMMITID=$version
COPY --from=builder /GoProject/src/dreba/_output/dreba /dreba
ENTRYPOINT ["/dreba"]

