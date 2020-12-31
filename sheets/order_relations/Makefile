order_relations.pdf: *.tex ../*.tex ../trademark.pdf
	pdflatex order_relations.tex
	make defs

make:
	make order_relations.pdf

open: order_relations.pdf
	open order_relations.pdf

defs: order_relations.tex
	defs order_relations.tex
