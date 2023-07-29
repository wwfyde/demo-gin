# See: https://go.dev/blog/docker
FROM golang:alpine
#MAINTAINER "Viakayn <wwfyde@gmail.com>"
LABEL maintainer="Viakayn <wwfyde@gmail.com>"

WORKDIR /demo-gin

COPY . .

# TODO  This layer need to be optimized
RUN #go mod download && go build .
RUN go build .

# Run the outyet command by default when the container starts.
# ENTRYPOINT /go/bin/outyet

EXPOSE 8000

CMD ["./demo-gin", "-port", "8000"]