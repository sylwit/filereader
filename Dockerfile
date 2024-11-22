# syntax=docker/dockerfile:1

FROM scratch

ARG TARGETARCH

COPY build/filereader-linux-${TARGETARCH} /filereader

CMD ["/filereader"]