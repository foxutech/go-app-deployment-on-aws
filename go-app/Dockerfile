FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
EXPOSE 8082 
RUN go build -o go-app . 
CMD ["/app/go-app"]
