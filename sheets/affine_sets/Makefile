# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make affine_sets.pdf

affine_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > affine_sets.tex

affine_sets.pdf: ../../*.tex ../../trademark.pdf *.tex affine_sets.tex
	# pdflatex --file-line-error -interaction=nonstopmode affine_sets.tex
	../../latexrun affine_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: affine_sets.pdf
	open affine_sets.pdf

o:
	make open

terms: affine_sets.tex
	bbk_sheet -mode ts affine_sets.tex

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
	rm -f affine_sets.*
	make reset
	bbk_sheet -mode mk > Makefile