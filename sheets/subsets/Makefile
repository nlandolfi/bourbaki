# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make subsets.pdf
	make graph.pdf

subsets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subsets.tex

subsets.pdf: ../../*.tex ../../trademark.pdf *.tex subsets.tex
	pdflatex subsets.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subsets.pdf
	open subsets.pdf

o:
	make open

defs: subsets.tex
	defs subsets.tex

spell:
	aspell -c subsets.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm subsets.*
	make reset
	bbk_sheet -mode mk > Makefile