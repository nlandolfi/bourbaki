# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_rings.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

matrix_rings.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_rings.tex

matrix_rings.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_rings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix_rings.tex
	../../latexrun matrix_rings.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_rings.pdf
	open matrix_rings.pdf

o:
	make open

terms: matrix_rings.tex
	bbk_sheet -mode ts matrix_rings.tex

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
	rm -f matrix_rings.*
	make reset
	bbk_sheet -mode mk > Makefile