# Dockerfile
FROM golang:1.22 AS builder
WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

# instalar ferramentas necessárias
RUN go install github.com/spf13/cobra-cli@latest || true
RUN go install github.com/golang/mock/mockgen@v1.5.0 || true

CMD ["tail", "-f", "/dev/null"]