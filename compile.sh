#!/usr/bin/env bash
root_path="api"
cur_path=`pwd`
api_proto_list=`ls ./$root_path/ -l | grep ^- | grep .proto | awk '{print $NF}'`
package_proto=`ls ./$root_path -l | grep ^d | awk '{print $NF}'`
package_file="message.proto"
googleapis="/home/zhangerlong/go/src/googleapis"
exec_cmd="protoc -I${cur_path}/${root_path} -I${googleapis} --go_out=plugins=grpc:. "
for row in $api_proto_list
do
   exec_cmd=${exec_cmd}" "${cur_path}/${root_path}/$row
done

for item in $package_proto
do
  exec_cmd=${exec_cmd}" "${cur_path}/${root_path}/${item}/${package_file}" "
done
$exec_cmd
echo $exec_cmd
#protoc -I/mnt/f/person_dir/framework/go-frame-examples/ -I$GOPATH/src/googleapis --go_out=plugins=grpc:. /mnt/f/person_dir/framework/go-frame-examples/api/demo.proto /mnt/f/person_dir/framework/go-frame-examples/api/api.proto