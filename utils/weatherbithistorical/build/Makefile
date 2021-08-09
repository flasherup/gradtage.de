build:
	env GOOS=linux GOARCH=amd64 go build -o weatherbithistory github.com/flasherup/gradtage.de/utils/weatherbithistorical


upload_to_cloud:
	pscp weatherbithistory flasherup@82.165.119.83:/home/flasherup/app/utils
	pscp configwbh.yml flasherup@82.165.119.83:/home/flasherup/app/utils
	pscp stations.csv flasherup@82.165.119.83:/home/flasherup/app/utils
