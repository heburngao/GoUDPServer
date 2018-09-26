path=$(dirname $0)
path=${path/\./$(pwd)}
echo $path

# /////////////////////////////////////////////////////////////////////////////
#
# 编译Protobuf协议
#
# /////////////////////////////////////////////////////////////////////////////

protoc --version

#protoc --go_out=$path/../ -I=$path OK.proto 
protoc --go_out=$path/../ -I=$path $path/helloworld.proto

#protoc --go_out=$path/../ -I=$path $path/PlayerInfo.proto
#protoc --go_out=$path/../ -I=$path $path/Rqst_CreateSelf.proto
#protoc --go_out=$path/../ -I=$path $path/Rqst_UpdateStatus.proto
#protoc --go_out=$path/../ -I=$path $path/Rqst_HeartBeating.proto

#protoc --go_out=$path/../ -I=$path $path/Rspn_CreateOthers.proto
#protoc --go_out=$path/../ -I=$path $path/Rspn_CreateSelf.proto
#protoc --go_out=$path/../ -I=$path $path/Rspn_UpdateStatus.proto
#protoc --go_out=$path/../ -I=$path $path/Rspn_HeartBeating.proto


 
protoc --go_out=$path/../msg -I=$path PlayerInfo.proto
protoc --go_out=$path/../msg -I=$path Rqst_CreateSelf.

#protoc --go_out=$path/../msg -I=$path Rqst_UpdateStatus.proto


#protoc --go_out=$path/../msg -I=$path Rspn_CreateOthers.proto
#protoc --go_out=$path/../msg -I=$path Rspn_CreateSelf.proto
#protoc --go_out=$path/../msg -I=$path Rspn_UpdateStatus.proto

protoc --go_out=$path/../msg -I=$path Rqst_HeartBeating.proto
protoc --go_out=$path/../msg -I=$path Rspn_HeartBeating.proto