setup_infrastructure:
	useradd --no-create-home --shell /bin/false gradtage
	mkdir /etc/gradtage
	chown gradtage:gradtage /etc/gradtage


setup_weatherbitsvc:
	sudo -u postgres createdb "hdd-weatherbit"
	cp ../weatherbitsvc /usr/local/bin/
	chmod 700 /usr/local/bin/weatherbitsvc
	chown gradtage:gradtage /usr/local/bin/weatherbitsvc
	cp weatherbitsvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/weatherbitsvc.yml
	cp weatherbitsvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start weatherbitsvc
	systemctl enable weatherbitsvc

setup_weatherbitupdatesvc:
	cp ../weatherbitupdatesvc /usr/local/bin/
	chmod 700 /usr/local/bin/weatherbitupdatesvc
	chown gradtage:gradtage /usr/local/bin/weatherbitupdatesvc
	cp weatherbitupdatesvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/weatherbitupdatesvc.yml
	cp weatherbitupdatesvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start weatherbitupdatesvc
	systemctl enable weatherbitupdatesvc

setup_daydegreesvc:
	cp ../daydegreesvc /usr/local/bin/
	chmod 700 /usr/local/bin/daydegreesvc
	chown gradtage:gradtage /usr/local/bin/daydegreesvc
	cp daydegreesvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/daydegreesvc.yml
	cp daydegreesvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start daydegreesvc
	systemctl enable daydegreesvc

setup_metricssvc:
	sudo -u postgres createdb "hdd-metrics"
	cp ../metricssvc /usr/local/bin/
	chmod 700 /usr/local/bin/metricssvc
	chown gradtage:gradtage /usr/local/bin/metricssvc
	cp metricssvc.yml /etc/gradtage/
	sudo chown gradtage:gradtage /etc/gradtage/metricssvc.yml
	cp metricssvc.service /etc/systemd/system/
	systemctl daemon-reload
	systemctl start metricssvc
	systemctl enable metricssvc