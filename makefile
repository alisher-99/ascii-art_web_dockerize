pusk:
	docker image rm ascii-taimas:ascii-docker-1.0
	docker build -t ascii-taimas:ascii-docker-1.0 .
	docker run -p 8181:8181 --rm ascii-taimas:ascii-docker-1.0