# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affine_transformations.pdf

affine_transformations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affine_transformations.tex

affine_transformations.pdf: ../../*.tex ../../trademark.pdf *.tex affine_transformations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode affine_transformations.tex
	../../latexrun affine_transformations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_transformations.pdf
	open affine_transformations.pdf

o:
	make open

terms: affine_transformations.tex
	bbk_sheet -mode ts affine_transformations.tex

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
	rm -f affine_transformations.*
	make reset
	bbk_sheet -mode mk > Makefile