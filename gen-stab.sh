docker run \
  --volume "$(pwd)/proto:/workspace" \
  --volume "$(pwd)/go:/go" \
  --workdir /workspace \
  bufbuild/buf generate
