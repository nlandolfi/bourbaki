# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_completeness.pdf
	make graph.pdf

metric_completeness.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_completeness.tex

metric_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completeness.tex
	pdflatex metric_completeness.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_completeness.pdf
	open metric_completeness.pdf

o:
	make open

defs: metric_completeness.tex
	defs metric_completeness.tex

spell:
	aspell -c metric_completeness.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm metric_completeness.*
	make reset
	bbk_sheet -mode mk > Makefile