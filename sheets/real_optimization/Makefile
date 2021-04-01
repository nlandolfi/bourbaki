# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_optimization.pdf
	make graph.pdf

real_optimization.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_optimization.tex

real_optimization.pdf: ../../*.tex ../../trademark.pdf *.tex real_optimization.tex
	pdflatex --file-line-error -interaction=nonstopmode real_optimization.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_optimization.pdf
	open real_optimization.pdf

o:
	make open

terms: real_optimization.tex
	bbk_sheet -mode ts real_optimization.tex

spell:
	aspell -c real_optimization.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_optimization.*
	make reset
	bbk_sheet -mode mk > Makefile