# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affine_hulls.pdf

affine_hulls.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affine_hulls.tex

affine_hulls.pdf: ../../*.tex ../../trademark.pdf *.tex affine_hulls.tex
	# pdflatex --file-line-error -interaction=nonstopmode affine_hulls.tex
	../../latexrun affine_hulls.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_hulls.pdf
	open affine_hulls.pdf

o:
	make open

terms: affine_hulls.tex
	bbk_sheet -mode ts affine_hulls.tex

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
	rm -f affine_hulls.*
	make reset
	bbk_sheet -mode mk > Makefile