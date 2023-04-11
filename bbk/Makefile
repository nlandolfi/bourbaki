compile:
	bbk compile_screen
local:
	bbk_serve -static-dir ./static -local-audit
deploy:
	gcloud beta run deploy bbk-www --project elos2-261103 --region us-west1 --platform managed --source . --project elos2-261103 --region us-west1
update:
	GOPRIVATE=github.com/nlandolfi go get -u github.com/nlandolfi/lit
