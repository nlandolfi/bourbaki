# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_function_graphs.pdf
	make graph.pdf

real_function_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_function_graphs.tex

real_function_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex real_function_graphs.tex
	pdflatex --file-line-error -interaction=nonstopmode real_function_graphs.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_function_graphs.pdf
	open real_function_graphs.pdf

o:
	make open

defs: real_function_graphs.tex
	defs real_function_graphs.tex

spell:
	aspell -c real_function_graphs.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm real_function_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile