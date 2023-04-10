# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make vector_subspaces.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

vector_subspaces.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > vector_subspaces.tex

vector_subspaces.pdf: ../../*.tex ../../trademark.pdf *.tex vector_subspaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode vector_subspaces.tex
	../../latexrun vector_subspaces.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vector_subspaces.pdf
	open vector_subspaces.pdf

o:
	make open

terms: vector_subspaces.tex
	bbk sheet -mode ts vector_subspaces.tex

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
	rm -f vector_subspaces.*
	make reset
	bbk sheet -mode mk > Makefile