FROM alpine

RUN apk add --no-cache bash && \
    apk add --no-cache ca-certificates

ADD app .

ENTRYPOINT ["./app"]
