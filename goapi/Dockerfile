# FROM golang:1.15.6
# RUN apt update \
#   && apt install -y vim
# ENV GO111MODULE on
# WORKDIR /go/src/work
# # install go tools
# RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
#   github.com/ramya-rao-a/go-outline \
#   github.com/nsf/gocode \
#   github.com/acroca/go-symbols \
#   github.com/fatih/gomodifytags \
#   github.com/josharian/impl \
#   github.com/haya14busa/goplay/cmd/goplay \
#   github.com/go-delve/delve/cmd/dlv \
#   golang.org/x/lint/golint \
#   golang.org/x/tools/gopls
# RUN go get -u github.com/cosmtrek/air


FROM golang:1.15.6 as builder
WORKDIR /go/src
COPY go.mod go.sum ./
RUN go mod download
COPY .  ./
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
  -o /go/bin/main \
  -ldflags '-s -w'
FROM scratch as runner
COPY --from=golang:1.15.6 /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /go/bin/main /app/main
ENTRYPOINT ["/app/main"]


