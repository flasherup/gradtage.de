setup_infrastructure:
	useradd --no-create-home --shell /bin/false gradtage
	mkdir /etc/gradtage
	chown gradtage:gradtage /etc/gradtage


setup_stations:
	sudo -u postgres createdb "hdd-stations"
	cp ../stationssvc /usr/local/bin/
	chmod 700 /usr/local/bin/stationssvc
	chown gradtage:gradtage /usr/local/bin/stationssvc
	cp stationssvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/stationssvc.yml
	cp stationssvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start stationssvc
	systemctl enable stationssvc

setup_user:
	sudo -u postgres createdb "hdd-user"
	cp ../usersvc /usr/local/bin/
	chmod 700 /usr/local/bin/usersvc
	chown gradtage:gradtage /usr/local/bin/usersvc
	cp usersvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/usersvc.yml
	cp usersvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start usersvc
	systemctl enable usersvc

setup_autocomplete:
	sudo -u postgres createdb "hdd-autocomplete"
	cp ../autocompletesvc /usr/local/bin/
	chmod 700 /usr/local/bin/autocompletesvc
	chown gradtage:gradtage /usr/local/bin/autocompletesvc
	cp autocompletesvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/autocompletesvc.yml
	cp autocompletesvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start autocompletesvc
	systemctl enable autocompletesvc

setup_api:
	cp ../apisvc /usr/local/bin/
	chmod 700 /usr/local/bin/apisvc
	chown gradtage:gradtage /usr/local/bin/apisvc
	cp apisvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/apisvc.yml
	cp apisvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start apisvc
	systemctl enable apisvc