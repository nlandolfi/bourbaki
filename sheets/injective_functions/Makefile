# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make injective_functions.pdf
	make graph.pdf

injective_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > injective_functions.tex

injective_functions.pdf: ../../*.tex ../../trademark.pdf *.tex injective_functions.tex
	pdflatex --file-line-error -interaction=nonstopmode injective_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: injective_functions.pdf
	open injective_functions.pdf

o:
	make open

terms: injective_functions.tex
	bbk_sheet -mode ts injective_functions.tex

spell:
	aspell -c injective_functions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f injective_functions.*
	make reset
	bbk_sheet -mode mk > Makefile