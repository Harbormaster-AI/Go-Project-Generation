FROM golang:1.16

RUN apt-get update --fix-missing && \
apt-get install -y git

RUN go get github.com/gorilla/mux
RUN go get -u gorm.io/gorm
RUN go get github.com/joho/godotenv
RUN go get -u gorm.io/driver/mysql

COPY . .

RUN cd src/go-demo

# RUN go mod init go-demo
# RUN go mod tidy

WORKDIR src/go-demo

RUN go build main.go

EXPOSE 8088

#Command to run the executable
CMD ["./go-demo"]