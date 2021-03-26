# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make outcome_vectors.pdf
	make graph.pdf

outcome_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > outcome_vectors.tex

outcome_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex outcome_vectors.tex
	pdflatex --file-line-error -interaction=nonstopmode outcome_vectors.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: outcome_vectors.pdf
	open outcome_vectors.pdf

o:
	make open

defs: outcome_vectors.tex
	defs outcome_vectors.tex

spell:
	aspell -c outcome_vectors.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm outcome_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile