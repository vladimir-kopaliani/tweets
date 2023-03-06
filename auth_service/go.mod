module github.com/vladimir-kopaliani/tweets/auth_service

go 1.20

require (
	github.com/99designs/gqlgen v0.17.25
	github.com/caarlos0/env/v6 v6.10.1
	github.com/golang-jwt/jwt/v5 v5.0.0-rc.1
	github.com/joho/godotenv v1.5.1
	github.com/rs/cors v1.8.3
	github.com/rs/zerolog v1.29.0
	github.com/vektah/gqlparser/v2 v2.5.1
	github.com/vladimir-kopaliani/tweets/user_service v0.0.0-00010101000000-000000000000
	golang.org/x/sync v0.1.0
	google.golang.org/grpc v1.53.0
)

replace github.com/vladimir-kopaliani/tweets/user_service => ./../user_service

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.24.4 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
