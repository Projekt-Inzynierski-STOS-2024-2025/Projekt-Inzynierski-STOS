FROM rust:latest
RUN apt-get update && apt-get install --no-install-recommends -y protobuf-compiler
RUN rustup toolchain install nightly
RUN rustup default nightly
COPY ./file-server /build
COPY ./messages.proto /
WORKDIR /build
RUN ["cargo", "build", "--release", "--verbose"]
ENTRYPOINT ["cargo", "run", "--release"]
