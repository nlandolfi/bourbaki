# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_probability_densities.pdf
	make graph.pdf

real_probability_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_probability_densities.tex

real_probability_densities.pdf: ../../*.tex ../../trademark.pdf *.tex real_probability_densities.tex
	pdflatex --file-line-error -interaction=nonstopmode real_probability_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_probability_densities.pdf
	open real_probability_densities.pdf

o:
	make open

terms: real_probability_densities.tex
	bbk_sheet -mode ts real_probability_densities.tex

spell:
	aspell -c real_probability_densities.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_probability_densities.*
	make reset
	bbk_sheet -mode mk > Makefile