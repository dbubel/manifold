FROM alpine:latest

WORKDIR /app

COPY ./dist/server .

EXPOSE 50051

CMD ["/app/server","serve"]
