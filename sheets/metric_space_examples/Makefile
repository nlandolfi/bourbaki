# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metric_space_examples.pdf
	make graph.pdf

metric_space_examples.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metric_space_examples.tex

metric_space_examples.pdf: ../../*.tex ../../trademark.pdf *.tex metric_space_examples.tex
	pdflatex --file-line-error -interaction=nonstopmode metric_space_examples.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metric_space_examples.pdf
	open metric_space_examples.pdf

o:
	make open

defs: metric_space_examples.tex
	defs metric_space_examples.tex

spell:
	aspell -c metric_space_examples.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm metric_space_examples.*
	make reset
	bbk_sheet -mode mk > Makefile