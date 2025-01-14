FROM --platform=$BUILDPLATFORM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Need to copy all files to include subpackages.
COPY . ./

# Build the static binary at /app/annai
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o annai

FROM scratch
WORKDIR /app
COPY --from=build /app/annai ./annai
EXPOSE 8080
ENTRYPOINT ["./annai"]