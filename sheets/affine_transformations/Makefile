# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affine_transformations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

affine_transformations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > affine_transformations.tex

affine_transformations.pdf: ../../*.tex ../../trademark.pdf *.tex affine_transformations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode affine_transformations.tex
	../../latexrun affine_transformations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_transformations.pdf
	open affine_transformations.pdf

o:
	make open

terms: affine_transformations.tex
	bbk sheet -mode ts affine_transformations.tex

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
	bbk sheet -mode mk > Makefile