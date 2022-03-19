FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /fiap-backend

EXPOSE 8080

CMD [ "/fiap-backend" ]