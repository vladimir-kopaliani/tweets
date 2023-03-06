FROM golang:1.20.1-alpine3.17 as builder
RUN apk --update --no-cache add tzdata ca-certificates protoc alpine-sdk && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
# RUN apk add gcompat # for CGO
WORKDIR /app
COPY auth_service auth_service
COPY user_service user_service
RUN cd user_service && \
    go generate ./... && \
    cd .. && \
    cd auth_service && \
    go generate ./... && \
    CGO_ENABLED=0 go build -o server ./cmd/app

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
# COPY configs/.env configs/.env
COPY --from=builder app/auth_service/server .
EXPOSE 8080
CMD ["./server"]
