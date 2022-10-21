FROM alpine:latest

WORKDIR /swagger

COPY api api
COPY swagger swagger
COPY main main

ENTRYPOINT ./main
