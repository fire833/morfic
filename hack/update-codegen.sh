#!/bin/bash

export GOPATH=$(go env GOPATH)

./update-deepcopy.sh
./update-clients.sh
./update-informers.sh
./update-listers.sh

# wait
