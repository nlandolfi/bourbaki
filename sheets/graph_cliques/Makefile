graph_cliques.tex: sheet.tex
	sheet -mode c -in sheet.tex > graph_cliques.tex

graph_cliques.pdf: ../../*.tex ../../trademark.pdf *.tex graph_cliques.tex
	pdflatex graph_cliques.tex
	make defs

make:
	make graph_cliques.pdf

open: graph_cliques.pdf
	open graph_cliques.pdf

o:
	make open

defs: graph_cliques.tex
	defs graph_cliques.tex

spell:
	aspell -c graph_cliques.tex

s:
	make spell