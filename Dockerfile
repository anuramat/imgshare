FROM golang as tool_builder
WORKDIR /tools
# build proto binaries
COPY go.* Makefile ./
RUN make dependencies

FROM bufbuild/buf AS generator
WORKDIR /generated
# glibc compatibility, make
RUN apk add libc6-compat make
# copy buf files, Makefile, proto files
COPY buf.* Makefile ./
COPY ./api ./api
# copy proto binaries
COPY --from=tool_builder /tools/bin/ ./bin/
# update buf dependencies and generate code/swagger
RUN make generate -o dependencies

FROM golang AS builder
WORKDIR /code
# download dependencies
COPY go.* ./
RUN go mod download
# copy code
COPY . ./
COPY --from=generator /generated ./
# build app binaries
RUN make build -o generate

FROM alpine as runner
# glibc compatibility
RUN apk add libc6-compat
WORKDIR /app
# copy binaries
COPY --from=builder /code/bin/imgshare* /code/bin/goose ./bin/
# copy swagger
COPY ./swagger_ui/ ./swagger_ui/
COPY --from=generator /generated/swagger/api/api.swagger.json ./swagger_ui/swagger.json
# migration stuff
COPY ./migrations ./migrations
COPY migrate.sh ./