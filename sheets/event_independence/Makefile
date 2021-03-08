# generated automatically by sheet.go

make:
	make event_independence.pdf

event_independence.tex: sheet.tex
	sheet -mode c -in sheet.tex > event_independence.tex

event_independence.pdf: ../../*.tex ../../trademark.pdf *.tex event_independence.tex
	pdflatex event_independence.tex
	make defs


open: event_independence.pdf
	open event_independence.pdf

o:
	make open

defs: event_independence.tex
	defs event_independence.tex

spell:
	aspell -c event_independence.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile