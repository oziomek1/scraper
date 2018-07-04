FROM golang:latest

ADD main.go /home/oziomek/project/scraper

WORKDIR /home/oziomek/projects/scraper


RUN go get .
RUN go build

CMD ["go", "main.go"]