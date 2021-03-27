# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make supremum_norm.pdf
	make graph.pdf

supremum_norm.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > supremum_norm.tex

supremum_norm.pdf: ../../*.tex ../../trademark.pdf *.tex supremum_norm.tex
	pdflatex --file-line-error -interaction=nonstopmode supremum_norm.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: supremum_norm.pdf
	open supremum_norm.pdf

o:
	make open

defs: supremum_norm.tex
	defs supremum_norm.tex

spell:
	aspell -c supremum_norm.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f supremum_norm.*
	make reset
	bbk_sheet -mode mk > Makefile