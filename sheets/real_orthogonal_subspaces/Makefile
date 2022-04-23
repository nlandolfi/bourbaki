# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_orthogonal_subspaces.pdf

real_orthogonal_subspaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_orthogonal_subspaces.tex

real_orthogonal_subspaces.pdf: ../../*.tex ../../trademark.pdf *.tex real_orthogonal_subspaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_orthogonal_subspaces.tex
	../../latexrun real_orthogonal_subspaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_orthogonal_subspaces.pdf
	open real_orthogonal_subspaces.pdf

o:
	make open

terms: real_orthogonal_subspaces.tex
	bbk_sheet -mode ts real_orthogonal_subspaces.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_orthogonal_subspaces.*
	make reset
	bbk_sheet -mode mk > Makefile