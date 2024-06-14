docker run \
  --volume "$(pwd)/proto/product:/workspace" \
  --volume "$(pwd)/go:/go" \
  --workdir /workspace \
  bufbuild/buf lint

docker run \
  --volume "$(pwd)/proto/product:/workspace" \
  --volume "$(pwd)/go:/go" \
  --workdir /workspace \
  bufbuild/buf generate
