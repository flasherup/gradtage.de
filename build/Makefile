build_weatherbit:
	env GOOS=linux GOARCH=amd64 go build -o weatherbitsvc github.com\flasherup\gradtage.de\weatherbitsvc\cmd\server


upload_weatherbit:
	pscp weatherbitsvc flasherup@212.227.214.163:/home/flasherup/app

update_weatherbit:
	ssh flasherup@212.227.214.163 "cd /home/flasherup/app && make update_weatherbitsvc $*"