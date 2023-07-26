#!/bin/bash

echo "::task~> Clean up & Build binaries files"
rm integration-test reveal 2>/dev/null
cd ../cmd/reveal
go build
mv reveal ../../integration-tests/reveal
cd ../../integration-tests
go build
echo "::done::"
echo "::task~> Run integration test"
./integration-tests
echo "::done::"
if [ $? -eq 0 ]
then
  exit 0
else
  exit 1
fi
