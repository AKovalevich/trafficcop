FROM scratch

COPY bin/linux-amd64/trafficcop /

CMD ["./trafficcop"]
