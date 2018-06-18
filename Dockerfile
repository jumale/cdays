FROM scratch

ENV PORT 8000
ENV DIAG_PORT 8001

EXPOSE $PORT
EXPOSE $DIAG_PORT

COPY ./bin/linux-amd64/cdays /

CMD ["/cdays"]
