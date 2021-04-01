# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make absolute_value.pdf
	make graph.pdf

absolute_value.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > absolute_value.tex

absolute_value.pdf: ../../*.tex ../../trademark.pdf *.tex absolute_value.tex
	pdflatex --file-line-error -interaction=nonstopmode absolute_value.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: absolute_value.pdf
	open absolute_value.pdf

o:
	make open

terms: absolute_value.tex
	bbk_sheet -mode ts absolute_value.tex

spell:
	aspell -c absolute_value.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f absolute_value.*
	make reset
	bbk_sheet -mode mk > Makefile