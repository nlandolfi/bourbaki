# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability_measures.pdf
	make graph.pdf

probability_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_measures.tex

probability_measures.pdf: ../../*.tex ../../trademark.pdf *.tex probability_measures.tex
	pdflatex --file-line-error -interaction=nonstopmode probability_measures.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_measures.pdf
	open probability_measures.pdf

o:
	make open

defs: probability_measures.tex
	defs probability_measures.tex

spell:
	aspell -c probability_measures.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm probability_measures.*
	make reset
	bbk_sheet -mode mk > Makefile