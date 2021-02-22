# generated automatically by sheet.go

make:
	make probability_events.pdf

probability_events.tex: sheet.tex
	sheet -mode c -in sheet.tex > probability_events.tex

probability_events.pdf: ../../*.tex ../../trademark.pdf *.tex probability_events.tex
	pdflatex probability_events.tex
	make defs


open: probability_events.pdf
	open probability_events.pdf

o:
	make open

defs: probability_events.tex
	defs probability_events.tex

spell:
	aspell -c probability_events.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile