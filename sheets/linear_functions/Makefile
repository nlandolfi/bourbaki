# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make linear_functions.pdf
	make graph.pdf

linear_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_functions.tex

linear_functions.pdf: ../../*.tex ../../trademark.pdf *.tex linear_functions.tex
	# pdflatex --file-line-error -interaction=nonstopmode linear_functions.tex
	../../latexrun linear_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_functions.pdf
	open linear_functions.pdf

o:
	make open

terms: linear_functions.tex
	bbk_sheet -mode ts linear_functions.tex

spell:
	aspell -c linear_functions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f linear_functions.*
	make reset
	bbk_sheet -mode mk > Makefile