FROM --platform=linux/amd64 golang:1.22.2-alpine as builder
WORKDIR /builder
ENV GO111MODULE=on CGO_ENABLED=0
COPY . .
RUN apk update && apk add --no-cache upx \
    && go build -ldflags "-extldflags '-static' -s -w" -o /builder/app ./src/cmd/app.go \
    && upx -9 /builder/app

FROM --platform=linux/amd64 alpine:3.16.0
WORKDIR /app
COPY --from=builder /builder/.env.example /app/.env
COPY --from=builder /builder/app /app/app