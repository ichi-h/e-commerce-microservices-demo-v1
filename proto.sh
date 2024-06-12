docker run \
  --volume "$(pwd)/proto:/workspace" \
  --workdir /workspace \
  bufbuild/buf generate
