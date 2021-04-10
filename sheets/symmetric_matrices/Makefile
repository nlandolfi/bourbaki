# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make symmetric_matrices.pdf
	make graph.pdf

symmetric_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > symmetric_matrices.tex

symmetric_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex symmetric_matrices.tex
	# pdflatex --file-line-error -interaction=nonstopmode symmetric_matrices.tex
	../../latexrun symmetric_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: symmetric_matrices.pdf
	open symmetric_matrices.pdf

o:
	make open

terms: symmetric_matrices.tex
	bbk_sheet -mode ts symmetric_matrices.tex

spell:
	aspell -c symmetric_matrices.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f symmetric_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile