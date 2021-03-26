# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability_distributions.pdf
	make graph.pdf

probability_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_distributions.tex

probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex probability_distributions.tex
	pdflatex --file-line-error -interaction=nonstopmode probability_distributions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_distributions.pdf
	open probability_distributions.pdf

o:
	make open

defs: probability_distributions.tex
	defs probability_distributions.tex

spell:
	aspell -c probability_distributions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm probability_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile