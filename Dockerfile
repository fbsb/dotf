FROM golang AS builder

ARG package=github.com/fbsb/dotf

COPY . /go/src/$package
RUN go install $package

FROM ubuntu

COPY --from=builder /go/bin/dotf /bin/dotf

WORKDIR /root

CMD ["/bin/dotf"]
