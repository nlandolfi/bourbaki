# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make multivariate_normal_maximum_likelihood.pdf
	make graph.pdf

multivariate_normal_maximum_likelihood.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_normal_maximum_likelihood.tex

multivariate_normal_maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_normal_maximum_likelihood.tex
	pdflatex --file-line-error -interaction=nonstopmode multivariate_normal_maximum_likelihood.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_normal_maximum_likelihood.pdf
	open multivariate_normal_maximum_likelihood.pdf

o:
	make open

defs: multivariate_normal_maximum_likelihood.tex
	defs multivariate_normal_maximum_likelihood.tex

spell:
	aspell -c multivariate_normal_maximum_likelihood.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm multivariate_normal_maximum_likelihood.*
	make reset
	bbk_sheet -mode mk > Makefile