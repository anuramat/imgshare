FROM golang as tool_builder
WORKDIR /tools
# build proto binaries
COPY go.* Makefile ./
RUN make dep

FROM bufbuild/buf AS generator
WORKDIR /generated
# glibc compatibility, make
RUN apk add libc6-compat make
# copy makefile, buf files, proto files
COPY Makefile ./
COPY buf* ./
COPY ./api ./api
# copy proto binaries
COPY --from=tool_builder /tools/bin/ ./bin/
# update buf dependencies and generate code/swagger
RUN make gen -o dep

FROM golang AS builder
WORKDIR /code
# download dependencies
COPY go.* ./
RUN go mod download
# copy code
COPY . ./
COPY --from=generator /generated ./internal
# build app binaries
RUN make build -o gen

FROM alpine as runner
# glibc compatibility
RUN apk add libc6-compat
WORKDIR /app
# copy binaries
COPY --from=builder /code/bin/bot /code/bin/db ./