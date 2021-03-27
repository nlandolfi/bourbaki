# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make commutative_operations.pdf
	make graph.pdf

commutative_operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > commutative_operations.tex

commutative_operations.pdf: ../../*.tex ../../trademark.pdf *.tex commutative_operations.tex
	pdflatex --file-line-error -interaction=nonstopmode commutative_operations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: commutative_operations.pdf
	open commutative_operations.pdf

o:
	make open

defs: commutative_operations.tex
	defs commutative_operations.tex

spell:
	aspell -c commutative_operations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f commutative_operations.*
	make reset
	bbk_sheet -mode mk > Makefile