FROM golang:alpine AS demoSvc
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go-paas/svc

COPY go.mod .
COPY go.sum .
ENV GOPROXY https://goproxy.cn
RUN go mod download

COPY . .

RUN go build -o go-paas-svc

FROM scratch

COPY ./conf /conf

COPY --from=demoSvc /go-paas/svc/go-paas-svc /

EXPOSE 8083
ENTRYPOINT ["/go-paas-svc", "conf/config.yaml"]