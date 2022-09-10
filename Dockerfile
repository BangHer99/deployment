FROM golang:1.18

# membuat direktori app
RUN mkdir /app

# set working directory /app
WORKDIR /app

COPY ./ /app

RUN go build -o ccid-fix

CMD ["./ccid-fix"]
