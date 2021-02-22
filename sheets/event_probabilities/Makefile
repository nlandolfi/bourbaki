# generated automatically by sheet.go

make:
	make event_probabilities.pdf

event_probabilities.tex: sheet.tex
	sheet -mode c -in sheet.tex > event_probabilities.tex

event_probabilities.pdf: ../../*.tex ../../trademark.pdf *.tex event_probabilities.tex
	pdflatex event_probabilities.tex
	make defs


open: event_probabilities.pdf
	open event_probabilities.pdf

o:
	make open

defs: event_probabilities.tex
	defs event_probabilities.tex

spell:
	aspell -c event_probabilities.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile