FROM golang:alpine
RUN mkdir /app
add ./test /app/
RUN adduser -S -D -H -h /app appuser
USER appuser
WORKDIR /app
CMD ["/app/test"]
