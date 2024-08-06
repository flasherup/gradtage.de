build_weatherbit:
	env GOOS=linux GOARCH=amd64 go build -o weatherbitsvc ./../weatherbitsvc/cmd/server

build_weatherbitupdate:
	env GOOS=linux GOARCH=amd64 go build -o weatherbitupdatesvc ./../weathrbitupdatesvc/cmd/server

build_api:
	env GOOS=linux GOARCH=amd64 go build -o apisvc ./../apisvc/cmd/server

build_users:
	env GOOS=linux GOARCH=amd64 go build -o usersvc ./../usersvc/cmd/server

build_autocomplete:
	env GOOS=linux GOARCH=amd64 go build -o autocompletesvc ./../autocompletesvc/cmd/server

build_daydegree:
	env GOOS=linux GOARCH=amd64 go build -o daydegreesvc ./../daydegreesvc/cmd/server


build_stations:
	env GOOS=linux GOARCH=amd64 go build -o stationssvc ./../stationssvc/cmd/server

build_metrics:
	env GOOS=linux GOARCH=amd64 go build -o metricssvc ./../metricssvc/cmd/server

build_all:
	make build_weatherbit
	make build_weatherbitupdate
	make build_api
	make build_users
	make build_autocomplete
	make build_stations
	make build_daydegree
	make build_metrics

upload_weatherbit:
	pscp weatherbitsvc flasherup@212.227.215.17:/home/flasherup/app

upload_weatherbitupdate:
	pscp weatherbitupdatesvc flasherup@212.227.215.17:/home/flasherup/app

upload_api:
	pscp apisvc flasherup@212.227.215.17:/home/flasherup/app

upload_users:
	scp usersvc flasherup@212.227.215.17:/home/flasherup/app

upload_autocomplete:
	scp autocompletesvc flasherup@212.227.215.17:/home/flasherup/app

upload_stations:
	scp stationssvc flasherup@212.227.215.17:/home/flasherup/app

update_weatherbit:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_weatherbitsvc $*"

update_weatherbitupdate:
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && make update_weatherbitupdatesvc $*"

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
	scp weatherbitsvc flasherup@82.165.119.83:/home/flasherup/app

upload_cloud_weatherbitupdate:
	scp weatherbitupdatesvc flasherup@82.165.119.83:/home/flasherup/app

upload_cloud_daydegree:
	scp daydegreesvc flasherup@82.165.119.83:/home/flasherup/app

upload_cloud_metrics:
	scp metricssvc flasherup@82.165.119.83:/home/flasherup/app


update_cloud_weatherbit:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_weatherbitsvc $*"

update_cloud_weatherbitupdate:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_weatherbitupdatesvc $*"

update_cloud_daydegree:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_daydegreesvc $*"


update_cloud_metrics:
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && make update_metricssvc $*"

deploy_cloud_weatherbit:
	make build_weatherbit
	make upload_cloud_weatherbit
	make update_cloud_weatherbit

deploy_cloud_weatherbitupdate:
	make build_weatherbitupdate
	make upload_cloud_weatherbitupdate
	make update_cloud_weatherbitupdate

deploy_cloud_daydegree:
	make build_daydegree
	make upload_cloud_daydegree
	make update_cloud_daydegree

deploy_cloud_metrics:
	make build_metrics
	make upload_cloud_metrics
	make update_cloud_metrics

deploy_cloud_all:
	make deploy_cloud_weatherbit
	make deploy_cloud_weatherbitupdate
	make deploy_cloud_daydegree

update_grpcs:
	make -C ../alertsvc/grpcalt/ regenerate
	make -C ../autocompletesvc/acrpc/ regenerate
	make -C ../daydegreesvc/ddgrpc/ regenerate
	make -C ../stationssvc/stsgrpc regenerate
	make -C ../usersvc/grpcusr regenerate
	make -C ../weatherbitsvc/weatherbitgrpc regenerate
	make -C ../weathrbitupdatesvc/wbugrpc regenerate
	make -C ../daydegreesvc/ddgrpc regenerate
	make -C ../metricssvc/mtrgrpc regenerate

setup_services:
	7z a -tzip setup.zip setup
	ssh flasherup@212.227.215.17 "rm -f /home/flasherup/app/setup.zip && rm -f -r /home/flasherup/app/setup"
	scp setup.zip flasherup@212.227.215.17:/home/flasherup/app
	rm setup.zip
	ssh flasherup@212.227.215.17 "cd /home/flasherup/app && unzip setup.zip"
	ssh flasherup@212.227.215.17 "find . -type f | xargs -n 5 touch && make clean"

setup_services_dc:
	7z a -tzip setup.zip setup_dc
	ssh flasherup@82.165.119.83 "rm -f /home/flasherup/app/setup.zip && rm -f -r /home/flasherup/app/setup"
	scp setup.zip flasherup@82.165.119.83:/home/flasherup/app
	ssh flasherup@82.165.119.83 "cd /home/flasherup/app && unzip setup.zip && mv setup_dc setup"
	rm setup.zip
	ssh flasherup@82.165.119.83 "find . -type f | xargs -n 5 touch && make clean"