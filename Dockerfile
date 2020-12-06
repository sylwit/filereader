FROM scratch

COPY build/filereader-linux-amd64 /filereader

CMD ["/filereader"]