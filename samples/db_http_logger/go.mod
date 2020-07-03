module github.com/spiral/cascade/samples/db_http_logger

go 1.14

require (
	github.com/NYTimes/gziphandler v1.1.1
	github.com/gorilla/mux v1.7.4
	github.com/rs/cors v1.7.0
	github.com/spiral/cascade v0.0.0-20200703110303-3631e2f5f56b
	go.etcd.io/bbolt v1.3.5
)

replace (
	github.com/spiral/cascade v0.0.0-20200703110303-3631e2f5f56b => /home/valery/Projects/opensource/spiral/cascade
)