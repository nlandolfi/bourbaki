# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make orthogonal_vectors.pdf
	make graph.pdf

orthogonal_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > orthogonal_vectors.tex

orthogonal_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_vectors.tex
	pdflatex --file-line-error -interaction=nonstopmode orthogonal_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orthogonal_vectors.pdf
	open orthogonal_vectors.pdf

o:
	make open

terms: orthogonal_vectors.tex
	bbk_sheet -mode ts orthogonal_vectors.tex

spell:
	aspell -c orthogonal_vectors.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f orthogonal_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile