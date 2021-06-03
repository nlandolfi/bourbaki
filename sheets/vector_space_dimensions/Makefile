# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make vector_space_dimensions.pdf

vector_space_dimensions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > vector_space_dimensions.tex

vector_space_dimensions.pdf: ../../*.tex ../../trademark.pdf *.tex vector_space_dimensions.tex
	# pdflatex --file-line-error -interaction=nonstopmode vector_space_dimensions.tex
	../../latexrun vector_space_dimensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vector_space_dimensions.pdf
	open vector_space_dimensions.pdf

o:
	make open

terms: vector_space_dimensions.tex
	bbk_sheet -mode ts vector_space_dimensions.tex

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
	bbk_sheet -mode mk > Makefile