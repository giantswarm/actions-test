FROM scratch
ADD ./actions-test /opt/bin/actions-test
ENTRYPOINT ["/opt/bin/actions-test"]
