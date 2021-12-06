compile:
	bbk_compile_screen
local:
	bbk_serve -static-dir ./static -local-audit
image:
	gcloud builds submit --tag gcr.io/elos2-261103/bbk-www --project elos2-261103
deploy:
	gcloud run deploy --image gcr.io/elos2-261103/bbk-www --platform managed --project elos2-261103
