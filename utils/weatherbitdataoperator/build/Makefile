build:
	env GOOS=linux GOARCH=amd64 go build -o weatherbitdataoperator github.com/flasherup/gradtage.de/utils/weatherbitdataoperator


upload_to_coloud:
	pscp weatherbitdataoperator flasherup@82.165.119.83:/home/flasherup/app/utils
	pscp wdoConfig.yml flasherup@82.165.119.83:/home/flasherup/app/utils
