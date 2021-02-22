# generated automatically by sheet.go

make:
	make conditional_event_probability.pdf

conditional_event_probability.tex: sheet.tex
	sheet -mode c -in sheet.tex > conditional_event_probability.tex

conditional_event_probability.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_event_probability.tex
	pdflatex conditional_event_probability.tex
	make defs


open: conditional_event_probability.pdf
	open conditional_event_probability.pdf

o:
	make open

defs: conditional_event_probability.tex
	defs conditional_event_probability.tex

spell:
	aspell -c conditional_event_probability.tex

s:
	make spell

remake:
	sheet -mode mk > Makefile