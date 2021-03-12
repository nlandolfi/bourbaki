# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_completion.pdf
	make graph.pdf

metric_completion.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_completion.tex

metric_completion.pdf: ../../*.tex ../../trademark.pdf *.tex metric_completion.tex
	pdflatex metric_completion.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_completion.pdf
	open metric_completion.pdf

o:
	make open

defs: metric_completion.tex
	defs metric_completion.tex

spell:
	aspell -c metric_completion.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm metric_completion.*
	make reset
	bbk_sheet -mode mk > Makefile