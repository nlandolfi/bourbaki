# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_inverses.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

matrix_inverses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_inverses.tex

matrix_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_inverses.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix_inverses.tex
	../../latexrun matrix_inverses.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_inverses.pdf
	open matrix_inverses.pdf

o:
	make open

terms: matrix_inverses.tex
	bbk_sheet -mode ts matrix_inverses.tex

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
	rm -f matrix_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile