FROM golang:1.23-bookworm AS build

WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/prometey-app ./cmd/server

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/prometey-app /prometey-app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/prometey-app"]
