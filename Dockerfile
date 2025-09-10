FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D -u 1001 -g 1001 sesamy

COPY sesamy /usr/bin/

USER sesamy
WORKDIR /home/sesamy

ENTRYPOINT ["sesamy"]
