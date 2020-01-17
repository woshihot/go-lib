module github.com/woshihot/go-lib

go 1.12

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/cheggaaa/pb v2.0.7+incompatible // indirect
	github.com/cheggaaa/pb/v3 v3.0.3
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/gorilla/mux v1.7.3
	github.com/hpcloud/tail v1.0.0
	github.com/json-iterator/go v1.1.9
	github.com/mitchellh/mapstructure v1.1.2
	github.com/rs/cors v1.7.0
	github.com/unrolled/render v1.0.1
	github.com/urfave/negroni v1.0.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	golang.org/x/sys v0.0.0-20191128015809-6d18c012aee9
	gopkg.in/VividCortex/ewma.v1 v1.1.1 // indirect
	gopkg.in/cheggaaa/pb.v2 v2.0.7 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200108215511-5d647ca15757
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20191209160850-c0dbc17a3553
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200116001909-b77594299b42
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20200108203644-89082a384178
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20191204190536-9bdfabe68543

)
