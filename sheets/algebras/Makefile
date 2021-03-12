# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make algebras.pdf
	make graph.pdf

algebras.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > algebras.tex

algebras.pdf: ../../*.tex ../../trademark.pdf *.tex algebras.tex
	pdflatex algebras.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: algebras.pdf
	open algebras.pdf

o:
	make open

defs: algebras.tex
	defs algebras.tex

spell:
	aspell -c algebras.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm algebras.*
	make reset
	bbk_sheet -mode mk > Makefile