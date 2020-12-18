product_metrics.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex product_metrics.tex
	make defs

make:
	make product_metrics.pdf

open: product_metrics.pdf
	open product_metrics.pdf

o:
	make open

defs: product_metrics.tex
	defs product_metrics.tex

spell:
	aspell -c product_metrics.tex

s:
	make spell
