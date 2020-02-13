
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.10

RUN git clone https://github.com/edenhill/librdkafka.git
WORKDIR librdkafka
RUN ./configure --prefix /usr
RUN make
RUN make install

RUN go get gopkg.in/confluentinc/confluent-kafka-go.v1/kafka
# Copy the local package files to the container's workspace.
ADD . /go/src/lynk/data-out-stream
# Build the ltl-service service inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install lynk/data-out-stream

# Run the ltl-service service by default when the container starts.
ENTRYPOINT /go/bin/data-out-stream

# Document that the service listens on port 3000.
EXPOSE 3000
