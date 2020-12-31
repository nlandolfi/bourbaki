graph_cliques.pdf: ../*.tex ../trademark.pdf *.tex
	pdflatex graph_cliques.tex
	make defs

make:
	make graph_cliques.pdf

open: graph_cliques.pdf
	open graph_cliques.pdf

defs: graph_cliques.tex
	defs graph_cliques.tex
