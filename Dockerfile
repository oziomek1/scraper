FROM golang:latest

ADD main.go /home/oziomek/project/scraper
ADD offer.go /home/oziomek/project/scraper
ADD parser.go /home/oziomek/project/scraper
ADD crawler.go /home/oziomek/project/scraper

WORKDIR /home/oziomek/projects/scraper


RUN go get .
RUN go build

CMD ["go", "main.go", "parser.go", "offer.go", "crawler.go"]
