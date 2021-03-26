# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make intervals.pdf
	make graph.pdf

intervals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > intervals.tex

intervals.pdf: ../../*.tex ../../trademark.pdf *.tex intervals.tex
	pdflatex --file-line-error -interaction=nonstopmode intervals.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: intervals.pdf
	open intervals.pdf

o:
	make open

defs: intervals.tex
	defs intervals.tex

spell:
	aspell -c intervals.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm intervals.*
	make reset
	bbk_sheet -mode mk > Makefile