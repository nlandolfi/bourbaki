# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cardinality.pdf
	make graph.pdf

cardinality.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cardinality.tex

cardinality.pdf: ../../*.tex ../../trademark.pdf *.tex cardinality.tex
	# pdflatex --file-line-error -interaction=nonstopmode cardinality.tex
	../../latexrun cardinality.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cardinality.pdf
	open cardinality.pdf

o:
	make open

terms: cardinality.tex
	bbk_sheet -mode ts cardinality.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f cardinality.*
	make reset
	bbk_sheet -mode mk > Makefile