# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make equivalence_relations.pdf
	make graph.pdf

equivalence_relations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > equivalence_relations.tex

equivalence_relations.pdf: ../../*.tex ../../trademark.pdf *.tex equivalence_relations.tex
	pdflatex --file-line-error -interaction=nonstopmode equivalence_relations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: equivalence_relations.pdf
	open equivalence_relations.pdf

o:
	make open

defs: equivalence_relations.tex
	defs equivalence_relations.tex

spell:
	aspell -c equivalence_relations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f equivalence_relations.*
	make reset
	bbk_sheet -mode mk > Makefile