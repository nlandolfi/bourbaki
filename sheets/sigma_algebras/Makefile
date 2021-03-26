# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make sigma_algebras.pdf
	make graph.pdf

sigma_algebras.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sigma_algebras.tex

sigma_algebras.pdf: ../../*.tex ../../trademark.pdf *.tex sigma_algebras.tex
	pdflatex --file-line-error -interaction=nonstopmode sigma_algebras.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sigma_algebras.pdf
	open sigma_algebras.pdf

o:
	make open

defs: sigma_algebras.tex
	defs sigma_algebras.tex

spell:
	aspell -c sigma_algebras.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm sigma_algebras.*
	make reset
	bbk_sheet -mode mk > Makefile