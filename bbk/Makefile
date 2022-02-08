compile:
	bbk_compile_screen
local:
	bbk_serve -static-dir ./static -local-audit
www:
	gcloud beta run deploy bbk-www --project elos2-261103 --region us-west1 --platform managed --source . --project elos2-261103 --region us-west1
