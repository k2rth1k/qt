FROM golang:latest
WORKDIR src
RUN BIN="/usr/local/bin" && \
    VERSION="1.0.0-rc2" && \
    BINARY_NAME="buf" && \
      curl -sSL \
        "https://github.com/bufbuild/buf/releases/download/v${VERSION}/${BINARY_NAME}-$(uname -s)-$(uname -m)" \
        -o "${BIN}/${BINARY_NAME}" && \
      chmod +x "${BIN}/${BINARY_NAME}" \
RUN git clone https://github.com/k2rth1k/qt.git
WORKDIR qt