all.pdf: ../*.tex ../trademark.pdf ../*/*.tex *.tex
	pdflatex all.tex
	pdflatex all.tex
	make defs

make:
	make all.pdf

open: all.pdf
	open all.pdf

o:
	make open

graph:
	cd .. && make open

defs: all.tex
	defs all.tex
