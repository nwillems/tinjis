FROM golang AS builder

ADD main.go /go/src/github.com/nwillems/tinjis/payment/main.go

# LDFlags from https://ops.tips/blog/dockerfile-golang/
RUN cd /go/src/github.com/nwillems/tinjis/payment/ && \
    CGO_ENABLED=0 go build -ldflags '-extldflags "-static"'

FROM scratch

ADD ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/nwillems/tinjis/payment/payment /payment

ENTRYPOINT [ "/payment" ]