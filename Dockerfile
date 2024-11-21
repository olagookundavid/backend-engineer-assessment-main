FROM alpine:latest

RUN mkdir /app

COPY bin/bookapp /app

CMD [ "/app/bookapp"]