# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make affine_set_dimensions.pdf
	make graph.pdf

affine_set_dimensions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affine_set_dimensions.tex

affine_set_dimensions.pdf: ../../*.tex ../../trademark.pdf *.tex affine_set_dimensions.tex
	# pdflatex --file-line-error -interaction=nonstopmode affine_set_dimensions.tex
	../../latexrun affine_set_dimensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_set_dimensions.pdf
	open affine_set_dimensions.pdf

o:
	make open

terms: affine_set_dimensions.tex
	bbk_sheet -mode ts affine_set_dimensions.tex

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
	rm -f affine_set_dimensions.*
	make reset
	bbk_sheet -mode mk > Makefile