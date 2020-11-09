module course

go 1.14

//解决 etcd clientv3 多个类undefined等问题
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
		github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/pkg/errors v0.9.1
	google.golang.org/protobuf v1.25.0
)