# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_continuity.pdf
	make graph.pdf

metric_continuity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_continuity.tex

metric_continuity.pdf: ../../*.tex ../../trademark.pdf *.tex metric_continuity.tex
	pdflatex --file-line-error -interaction=nonstopmode metric_continuity.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_continuity.pdf
	open metric_continuity.pdf

o:
	make open

defs: metric_continuity.tex
	defs metric_continuity.tex

spell:
	aspell -c metric_continuity.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm metric_continuity.*
	make reset
	bbk_sheet -mode mk > Makefile