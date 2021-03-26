# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make norm_metrics.pdf
	make graph.pdf

norm_metrics.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > norm_metrics.tex

norm_metrics.pdf: ../../*.tex ../../trademark.pdf *.tex norm_metrics.tex
	pdflatex --file-line-error -interaction=nonstopmode norm_metrics.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norm_metrics.pdf
	open norm_metrics.pdf

o:
	make open

defs: norm_metrics.tex
	defs norm_metrics.tex

spell:
	aspell -c norm_metrics.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm norm_metrics.*
	make reset
	bbk_sheet -mode mk > Makefile