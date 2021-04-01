# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make positive_definite_matrices.pdf
	make graph.pdf

positive_definite_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > positive_definite_matrices.tex

positive_definite_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex positive_definite_matrices.tex
	pdflatex --file-line-error -interaction=nonstopmode positive_definite_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: positive_definite_matrices.pdf
	open positive_definite_matrices.pdf

o:
	make open

terms: positive_definite_matrices.tex
	bbk_sheet -mode ts positive_definite_matrices.tex

spell:
	aspell -c positive_definite_matrices.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f positive_definite_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile