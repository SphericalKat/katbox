FROM alpine

WORKDIR /app

COPY dist/katbox /usr/bin

ENTRYPOINT [ "katbox" ]