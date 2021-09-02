build_weatherbit:
	env GOOS=linux GOARCH=amd64 go build -o weatherbitsvc github.com\flasherup\gradtage.de\weatherbitsvc\cmd\server

build_api:
	env GOOS=linux GOARCH=amd64 go build -o apisvc github.com\flasherup\gradtage.de\apisvc\cmd\server

build_users:
	env GOOS=linux GOARCH=amd64 go build -o usersvc github.com\flasherup\gradtage.de\usersvc\cmd\server

build_autocomplete:
	env GOOS=linux GOARCH=amd64 go build -o autocompletesvc github.com\flasherup\gradtage.de\autocompletesvc\cmd\server

build_daydegree:
	env GOOS=linux GOARCH=amd64 go build -o daydegreesvc github.com\flasherup\gradtage.de\daydegreesvc\cmd\server

build_stations:
	env GOOS=linux GOARCH=amd64 go build -o stationssvc github.com\flasherup\gradtage.de\stationssvc\cmd\server

build_all:
	make build_weatherbit
	make build_api
	make build_users
	make build_autocomplete
	make build_stations

upload_weatherbit:
	pscp weatherbitsvc flasherup@212.227.215.17:/home/flasherup/app

upload_api:
	pscp apisvc flasherup@212.227.215.17:/home/flasherup/app

upload_users:
	pscp usersvc flasherup@212.227.215.17:/home/flasherup/app

upload_autocomplete:
	pscp autocompletesvc flasherup@212.227.215.17:/home/flasherup/app

upload_stations:
	pscp stationssvc flasherup@212.227.215.17:/home/flasherup/app


update_weatherbit:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_weatherbitsvc $*"

update_api:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_apisvc $*"

update_users:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_usersvc $*"

update_autocomplete:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_autocompletesvc $*"

update_stations:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_stationssvc $*"


deploy_autocomplete:
	make build_autocomplete
	make upload_autocomplete
	make update_autocomplete

deploy_users:
	make build_users
	make upload_users
	make update_users

deploy_stations:
	make build_stations
	make upload_stations
	make update_stations

deploy_api:
	make build_api
	make upload_api
	make update_api

deploy_weatherbit:
	make build_weatherbit
	make upload_weatherbit
	make update_weatherbit

deploy_all:
	make deploy_autocomplete
	make deploy_users
	make deploy_stations
	make deploy_api
	#make deploy_weatherbit

upload_cloud_weatherbit:
	pscp weatherbitsvc flasherup@82.165.119.83:/home/flasherup/app

upload_cloud_daydegree:
	pscp daydegreesvc flasherup@82.165.119.83:/home/flasherup/app


update_cloud_weatherbit:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_weatherbitsvc $*"

update_cloud_daydegree:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_daydegreesvc $*"

deploy_cloud_weatherbit:
	make build_weatherbit
	make upload_cloud_weatherbit
	make update_cloud_weatherbit

deploy_cloud_daydegree:
	make build_daydegree
	make upload_cloud_daydegree
	make update_cloud_daydegree

update_grpcs:
	make -C ../autocompletesvc/acrpc/ regenerate
	make -C ../daydegreesvc/ddgrpc/ regenerate
	make -C ../stationssvc/stsgrpc regenerate
	make -C ../usersvc/grpcusr regenerate
	make -C ../weatherbitsvc/weatherbitgrpc regenerate

setup_services:
	7z a -tzip setup.zip setup
	ssh flasherup@212.227.215.17 "rm -f /home/flasherup/app/setup.zip && rm -f -r /home/flasherup/app/setup"
	pscp setup.zip flasherup@212.227.215.17:/home/flasherup/app
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && unzip setup.zip"
	ssh flasherup@212.227.215.17 "find . -type f | xargs -n 5 touch && make clean"
