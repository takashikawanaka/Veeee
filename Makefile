develop:
	air -c .\.air.toml

develop_go:
	go build -o ./tmp/main.exe ./...

develop_solid:
	cd web && corepack pnpm build

windows:
	go build -a -tags netgo -installsuffix netgo --ldflags '-extldflags "-static"' .\...

linux:
	$$env:GOOS="linux"
	go build -a -tags netgo  .\...
