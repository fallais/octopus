FROM golang:latest as builder
WORKDIR /go/src/octopus
ADD . /go/src/octopus
RUN go get -d -v ./...
RUN go build -o /go/bin/octopus

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/octopus /
ENTRYPOINT [ "/octopus" ]
CMD [ "--help" ]