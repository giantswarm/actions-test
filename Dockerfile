FROM scratch
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY actions-test /opt/bin/actions-test
ENTRYPOINT ["/opt/bin/actions-test"]
