FROM golang:1.16-alpine

RUN apk update && apk add --no-cache gcc
RUN apk add --update gcc musl-dev

WORKDIR /app
COPY . ./
RUN go build -o /docker-forum

EXPOSE 8080

CMD [ "/docker-forum" ]