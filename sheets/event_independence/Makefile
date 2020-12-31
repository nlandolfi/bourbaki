event_independence.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex event_independence.tex
	make defs

make:
	make event_independence.pdf

open: event_independence.pdf
	open event_independence.pdf

o:
	make open

defs: event_independence.tex
	defs event_independence.tex
