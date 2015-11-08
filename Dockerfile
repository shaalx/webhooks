FROM golang

# env
WORKDIR /app/go/
ENV GOPATH /app/go
ADD . /app/go/

# build
RUN go get github.com/shaalx/webhooks
RUN go build -o webhooks
EXPOSE 80

# run
CMD ["/app/go/webhooks"]
