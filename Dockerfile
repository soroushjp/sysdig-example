FROM instrumentisto/glide:latest

ADD . /go/src/github.com/soroushjp/sysdig-example
WORKDIR /go/src/github.com/soroushjp/sysdig-example
RUN glide install

EXPOSE 80

ENTRYPOINT ["go", "run", "main.go"]
