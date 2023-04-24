FROM golang:alpine AS demoPodApi
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go-paas/podApi

COPY go.mod .
COPY go.sum .
ENV GOPROXY https://goproxy.cn
RUN go mod download

COPY . .

RUN go build -o go-paas-podApi

FROM scratch

COPY ./conf /conf

COPY --from=demoPodApi /go-paas/podApi/go-paas-podApi /

EXPOSE 8082
ENTRYPOINT ["/go-paas-podApi", "conf/config.yaml"]