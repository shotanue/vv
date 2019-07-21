FROM znly/protoc:0.4.0

ENV PROJECT_ROOT=/work

WORKDIR /work
COPY . .

RUN mkdir -p go/protocol js/protocol

RUN /bin/sh ./scripts/compile_proto.sh golang
RUN /bin/sh ./scripts/compile_proto.sh js

# znly/protoc:0.4.0 の ENTRYPOINTは上書き
ENTRYPOINT ["sh", "./scripts/compile_proto.sh"]
CMD [""]