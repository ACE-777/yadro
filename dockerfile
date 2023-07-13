FROM golang:latest

WORKDIR /app

COPY . .

RUN go build cmd/yadro/main.go

CMD ["./main", "cmd/yadro/testFiles/test_file_1.txt"]