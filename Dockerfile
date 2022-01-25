# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

RUN mkdir /build
WORKDIR /build

# Copy the local package files to the container's workspace.
#ADD . /modules

RUN export GO111MODULE=on

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/jbabinec1/hello-world-test-service

RUN cd /build && git clone https://github.com/jbabinec1/hello-world-test-service.git

##RUN cd /build/hello-world-test-service && go build
RUN cd /build/hello-world-test-service

# Run the outyet command by default when the container starts.
#CMD ["go", "run", "main.go"]

# Document that the service listens on port 8080.
EXPOSE 6666

CMD ["go", "run", "main.go"]

#ENTRYPOINT [ "[/hello-world-test-service/hello-world-test-service]" ]