# generated automatically by sheet.go

make:
	make sigma_algebra_event_independence.pdf

sigma_algebra_event_independence.tex: sheet.tex
	sheet -mode c -in sheet.tex > sigma_algebra_event_independence.tex

sigma_algebra_event_independence.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebra_event_independence.tex
	pdflatex sigma_algebra_event_independence.tex
	make defs


open: sigma_algebra_event_independence.pdf
	open sigma_algebra_event_independence.pdf

o:
	make open

defs: sigma_algebra_event_independence.tex
	defs sigma_algebra_event_independence.tex

spell:
	aspell -c sigma_algebra_event_independence.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile