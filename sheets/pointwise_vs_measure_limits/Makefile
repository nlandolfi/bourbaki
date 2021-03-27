# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make pointwise_vs_measure_limits.pdf
	make graph.pdf

pointwise_vs_measure_limits.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > pointwise_vs_measure_limits.tex

pointwise_vs_measure_limits.pdf: ../../*.tex ../../trademark.pdf *.tex pointwise_vs_measure_limits.tex
	pdflatex --file-line-error -interaction=nonstopmode pointwise_vs_measure_limits.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: pointwise_vs_measure_limits.pdf
	open pointwise_vs_measure_limits.pdf

o:
	make open

defs: pointwise_vs_measure_limits.tex
	defs pointwise_vs_measure_limits.tex

spell:
	aspell -c pointwise_vs_measure_limits.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f pointwise_vs_measure_limits.*
	make reset
	bbk_sheet -mode mk > Makefile