# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability_densities.pdf
	make graph.pdf

probability_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_densities.tex

probability_densities.pdf: ../../*.tex ../../trademark.pdf *.tex probability_densities.tex
	pdflatex --file-line-error -interaction=nonstopmode probability_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_densities.pdf
	open probability_densities.pdf

o:
	make open

terms: probability_densities.tex
	bbk_sheet -mode ts probability_densities.tex

spell:
	aspell -c probability_densities.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f probability_densities.*
	make reset
	bbk_sheet -mode mk > Makefile