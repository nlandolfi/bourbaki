# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make multivariate_normal_conditionals.pdf
	make graph.pdf

multivariate_normal_conditionals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > multivariate_normal_conditionals.tex

multivariate_normal_conditionals.pdf: ../../*.tex ../../trademark.pdf *.tex multivariate_normal_conditionals.tex
	pdflatex --file-line-error -interaction=nonstopmode multivariate_normal_conditionals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: multivariate_normal_conditionals.pdf
	open multivariate_normal_conditionals.pdf

o:
	make open

terms: multivariate_normal_conditionals.tex
	bbk_sheet -mode ts multivariate_normal_conditionals.tex

spell:
	aspell -c multivariate_normal_conditionals.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f multivariate_normal_conditionals.*
	make reset
	bbk_sheet -mode mk > Makefile