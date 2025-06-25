module github.com/jackwrfuller/lagoon-go-client

go 1.23

require (
	github.com/adrg/xdg v0.5.3
	github.com/cxmcc/unixsums v0.0.0-20131125091133-89564297d82f
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/guregu/null v4.0.0+incompatible
	github.com/hashicorp/go-version v1.7.0
	github.com/machinebox/graphql v0.2.2
	github.com/robfig/cron/v3 v3.0.1
	github.com/uselagoon/machinery v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.32.0
	golang.org/x/oauth2 v0.25.0
	golang.org/x/term v0.28.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.29.0 // indirect
)

replace github.com/uselagoon/machinery => github.com/jackwrfuller/lagoon-go-client v0.0.0-20250625002020-089aac41d568
