module github.com/rhythmicsoul/nginx-mgmt

go 1.17

require (
	github.com/rhythmicsoul/nginx-mgmt/proto/controller v0.0.0-20220301170035-1ad0d1f883e1
	google.golang.org/grpc v1.44.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)

replace github.com/rhythmicsoul/nginx-mgmt/proto/controller => ./proto/controller
