FROM golang:1.16-alpine3.12 AS BUILD-ENV

RUN apk update && apk add git curl build-base

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
# -o denotes where to put executable
RUN go build -o /go/bin/server .

# Production image
FROM alpine:3.12

# Create Non Privileged user
RUN addgroup --gid 101 app && \
    adduser -S --uid 101 --ingroup app app

RUN mkdir -p /opt/server && chown -R app:app /opt/server
VOLUME /opt/server
COPY config.yml /opt/server

# Run as Non Privileged user
USER app

COPY --from=BUILD-ENV /go/bin/server /go/bin/server

CMD ["/go/bin/server"]