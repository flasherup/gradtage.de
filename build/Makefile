build_weatherbit:
	env GOOS=linux GOARCH=amd64 go build -o weatherbitsvc github.com\flasherup\gradtage.de\weatherbitsvc\cmd\server

build_api:
	env GOOS=linux GOARCH=amd64 go build -o apisvc github.com\flasherup\gradtage.de\apisvc\cmd\server

build_users:
	env GOOS=linux GOARCH=amd64 go build -o usersvc github.com\flasherup\gradtage.de\usersvc\cmd\server


upload_weatherbit:
	pscp weatherbitsvc flasherup@212.227.214.163:/home/flasherup/app

upload_api:
	pscp apisvc flasherup@212.227.214.163:/home/flasherup/app

upload_users:
	pscp usersvc flasherup@212.227.214.163:/home/flasherup/app


update_weatherbit:
	ssh flasherup@212.227.214.163 "cd /home/flasherup/app && make update_weatherbitsvc $*"

update_api:
	ssh flasherup@212.227.214.163 "cd /home/flasherup/app && make update_apisvc $*"

update_users:
	ssh flasherup@212.227.214.163 "cd /home/flasherup/app && make update_usersvc $*"


upload_cloud_weatherbit:
	pscp weatherbitsvc flasherup@82.165.119.83:/home/flasherup/app

update_cloud_weatherbit:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_weatherbitsvc $*"