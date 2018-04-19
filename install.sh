set -e
export GOPATH=$(cd `dirname $0`; pwd)
go install golang/www
