# /////////////////////////////////////////////////////////////////////////////
#
# 编译Protobuf2协议
#
# /////////////////////////////////////////////////////////////////////////////

protoc --version

protoc --go_out=../msg  *.proto