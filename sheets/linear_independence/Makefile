# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make linear_independence.pdf
	make graph.pdf

linear_independence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_independence.tex

linear_independence.pdf: ../../*.tex ../../trademark.pdf *.tex linear_independence.tex
	# pdflatex --file-line-error -interaction=nonstopmode linear_independence.tex
	../../latexrun linear_independence.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_independence.pdf
	open linear_independence.pdf

o:
	make open

terms: linear_independence.tex
	bbk_sheet -mode ts linear_independence.tex

spell:
	aspell -c linear_independence.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f linear_independence.*
	make reset
	bbk_sheet -mode mk > Makefile