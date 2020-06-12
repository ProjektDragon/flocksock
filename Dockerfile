# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-flocksock"
LABEL REPO="https://github.com/mpmsimo/flocksock"

ENV PROJPATH=/go/src/github.com/mpmsimo/flocksock

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/mpmsimo/flocksock
WORKDIR /go/src/github.com/mpmsimo/flocksock

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/mpmsimo/flocksock"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/flocksock/bin

WORKDIR /opt/flocksock/bin

COPY --from=build-stage /go/src/github.com/mpmsimo/flocksock/bin/flocksock /opt/flocksock/bin/
RUN chmod +x /opt/flocksock/bin/flocksock

# Create appuser
RUN adduser -D -g '' flocksock
USER flocksock

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/flocksock/bin/flocksock"]
