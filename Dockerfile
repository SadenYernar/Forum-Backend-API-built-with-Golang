FROM golang:1.19-alpine AS build 
LABEL stage=forum_builder

WORKDIR /build 
COPY . /build 
ENV CGO_ENABLED=1
RUN apk add gcc musl-dev && \
    go build -o forum cmd/main.go

FROM alpine 
COPY --from=build /build/forum /app/cmd/forum  
COPY --from=build /build/internal/template/ /app/internal/template/

ENV SERVER_ADDR="0.0.0.0"

CMD mkdir -p /app/db/SQLite3 && cd /app && ./cmd/forum 