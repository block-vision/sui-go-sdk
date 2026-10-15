#!/usr/bin/env bash
# Regenerate the Go stubs for the Sui gRPC v2 protos into pb/sui/rpc/v2.
#
# Prerequisites:
#   brew install protobuf
#   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.10
#   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
#
# Proto source: the `proto` submodule (https://github.com/MystenLabs/sui-apis),
# proto/sui/rpc/v2. To bump: `git -C proto checkout <commit>`, re-run this
# script, and commit the submodule pointer together with the regenerated files.
set -euo pipefail
cd "$(dirname "$0")/.."

MODULE=github.com/block-vision/sui-go-sdk
GO_PKG="$MODULE/pb/sui/rpc/v2"

FILES=$(cd proto/proto && find sui/rpc/v2 -name '*.proto' | sort)

M_FLAGS=()
for f in $FILES; do
  M_FLAGS+=("--go_opt=M$f=$GO_PKG" "--go-grpc_opt=M$f=$GO_PKG")
done
# google/rpc/status.proto is vendored only for protoc's import resolution; its
# Go code comes from the published genproto package, so map instead of generate.
M_FLAGS+=(
  "--go_opt=Mgoogle/rpc/status.proto=google.golang.org/genproto/googleapis/rpc/status"
  "--go-grpc_opt=Mgoogle/rpc/status.proto=google.golang.org/genproto/googleapis/rpc/status"
)

protoc -I proto/proto \
  --go_out=. --go_opt=module=$MODULE \
  --go-grpc_out=. --go-grpc_opt=module=$MODULE \
  "${M_FLAGS[@]}" \
  $FILES

echo "generated $(ls pb/sui/rpc/v2/*.go | wc -l | tr -d ' ') files into pb/sui/rpc/v2"
