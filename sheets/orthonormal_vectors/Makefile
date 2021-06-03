# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make orthonormal_vectors.pdf

orthonormal_vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > orthonormal_vectors.tex

orthonormal_vectors.pdf: ../../*.tex ../../trademark.pdf *.tex orthonormal_vectors.tex
	# pdflatex --file-line-error -interaction=nonstopmode orthonormal_vectors.tex
	../../latexrun orthonormal_vectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orthonormal_vectors.pdf
	open orthonormal_vectors.pdf

o:
	make open

terms: orthonormal_vectors.tex
	bbk_sheet -mode ts orthonormal_vectors.tex

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
	rm -f orthonormal_vectors.*
	make reset
	bbk_sheet -mode mk > Makefile