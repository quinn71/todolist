FROM golang:1.22.1-alpine3.19
WORKDIR /todolist
COPY . .
RUN go build -o todolist

CMD ["./todolist"]

EXPOSE 8020
