FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o main ./infrastructure/cmd/main.go

EXPOSE 9000

CMD ["./main"]

# docker build -t [username]/[image_name]:[version] .
# docker push [username]/[image_name]:[version]