# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make vector_space_dimensions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

vector_space_dimensions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > vector_space_dimensions.tex

vector_space_dimensions.pdf: ../../*.tex ../../trademark.pdf *.tex vector_space_dimensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode vector_space_dimensions.tex
	../../latexrun vector_space_dimensions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vector_space_dimensions.pdf
	open vector_space_dimensions.pdf

o:
	make open

terms: vector_space_dimensions.tex
	bbk sheet -mode ts vector_space_dimensions.tex

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
	rm -f vector_space_dimensions.*
	make reset
	bbk sheet -mode mk > Makefile