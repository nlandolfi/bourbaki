# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make multivariate_real_densities.pdf
	make graph.pdf

multivariate_real_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_real_densities.tex

multivariate_real_densities.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_real_densities.tex
	pdflatex --file-line-error -interaction=nonstopmode multivariate_real_densities.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_real_densities.pdf
	open multivariate_real_densities.pdf

o:
	make open

defs: multivariate_real_densities.tex
	defs multivariate_real_densities.tex

spell:
	aspell -c multivariate_real_densities.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm multivariate_real_densities.*
	make reset
	bbk_sheet -mode mk > Makefile