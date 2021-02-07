# Build Stage
FROM sonic3021/golang-buildimage:1.13 AS build-stage

LABEL app="build-sonic-trans-post"
LABEL REPO="https://github.com/sonic3021/sonic-trans-post"

ENV PROJPATH=/go/src/github.com/sonic3021/sonic-trans-post

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/sonic3021/sonic-trans-post
WORKDIR /go/src/github.com/sonic3021/sonic-trans-post

RUN make build-alpine

# Final Stage
FROM sonic3021/golang-base-image:lastest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/sonic3021/sonic-trans-post"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/sonic-trans-post/bin

WORKDIR /opt/sonic-trans-post/bin

COPY --from=build-stage /go/src/github.com/sonic3021/sonic-trans-post/bin/sonic-trans-post /opt/sonic-trans-post/bin/
RUN chmod +x /opt/sonic-trans-post/bin/sonic-trans-post

# Create appuser
RUN adduser -D -g '' sonic-trans-post
USER sonic-trans-post

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/sonic-trans-post/bin/sonic-trans-post"]
