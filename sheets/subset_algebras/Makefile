# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make subset_algebras.pdf
	make graph.pdf

subset_algebras.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subset_algebras.tex

subset_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex subset_algebras.tex
	pdflatex --file-line-error -interaction=nonstopmode subset_algebras.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subset_algebras.pdf
	open subset_algebras.pdf

o:
	make open

defs: subset_algebras.tex
	defs subset_algebras.tex

spell:
	aspell -c subset_algebras.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm subset_algebras.*
	make reset
	bbk_sheet -mode mk > Makefile