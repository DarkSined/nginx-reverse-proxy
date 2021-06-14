FROM golang:1.16.5-alpine3.13 as builder
LABEL stage=builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:3.13
COPY --from=builder /build/app .

ENTRYPOINT [ "./app" ]