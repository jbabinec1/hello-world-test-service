# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

#RUN mkdir /build
WORKDIR /go/src/hello-world-test-service

ADD . .


RUN export GO111MODULE=auto


RUN go get github.com/jbabinec1/hello-world-test-service


RUN go build


# Document that the service listens on port 8080.
EXPOSE 6666:6666

 #run main.go when container starts
CMD ["go", "run", "main.go"]

