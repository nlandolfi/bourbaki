# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make orthogonal_matrices.pdf

orthogonal_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > orthogonal_matrices.tex

orthogonal_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode orthogonal_matrices.tex
	../../latexrun orthogonal_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orthogonal_matrices.pdf
	open orthogonal_matrices.pdf

o:
	make open

terms: orthogonal_matrices.tex
	bbk_sheet -mode ts orthogonal_matrices.tex

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
	rm -f orthogonal_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile