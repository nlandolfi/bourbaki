compile:
	bbk_compile_screen

index.html:
	bbk_compile_screen

o: index.html
	open index.html

local:
	go run server.go -static-dir ./static
image:
	gcloud builds submit --tag gcr.io/elos2-261103/bbk-www
deploy:
	gcloud run deploy --image gcr.io/elos2-261103/bbk-www --platform managed
