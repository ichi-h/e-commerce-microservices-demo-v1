#!/usr/bin/zsh

function gen_index_ts() {
  # 現在のディレクトリを取得
  current_dir=$(pwd)

  # 引数で取得したディレクトリへ移動
  dir=$1
  cd $dir

  files=( $(ls -p | grep -v /) )

  # index.tsを初期化
  echo "" > index.ts

  # 各ファイルに対して
  for file in "${files[@]}"
  do
    if [ $file = "index.ts" ]; then
      continue
    fi
    # ファイル名から拡張子を除去
    filename=$(basename "$file" .ts)
    # index.tsに書き込み
    echo "export * from './$filename';" >> index.ts
  done
}

docker run \
  --volume "$(pwd)/proto/product:/workspace" \
  --workdir /workspace \
  bufbuild/buf lint

docker run \
  --volume "$(pwd)/proto/product:/workspace" \
  --volume "$(pwd)/go:/go" \
  --volume "$(pwd)/node:/node" \
  --workdir /workspace \
  bufbuild/buf generate

connect_dirs=('product-connect')

for connect_dir in "${connect_dirs[@]}"
do
  sudo chmod -R 777 ./node/packages/$connect_dir/api
  gen_index_ts ./node/packages/$connect_dir/api/v1
done
