# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make set_operations.pdf
	make graph.pdf

set_operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_operations.tex

set_operations.pdf: ../../*.tex ../../trademark.pdf *.tex set_operations.tex
	pdflatex --file-line-error -interaction=nonstopmode set_operations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_operations.pdf
	open set_operations.pdf

o:
	make open

terms: set_operations.tex
	bbk_sheet -mode ts set_operations.tex

spell:
	aspell -c set_operations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f set_operations.*
	make reset
	bbk_sheet -mode mk > Makefile