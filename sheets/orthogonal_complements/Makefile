# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make orthogonal_complements.pdf
	make graph.pdf

orthogonal_complements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > orthogonal_complements.tex

orthogonal_complements.pdf: ../../*.tex ../../trademark.pdf *.tex orthogonal_complements.tex
	pdflatex --file-line-error -interaction=nonstopmode orthogonal_complements.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orthogonal_complements.pdf
	open orthogonal_complements.pdf

o:
	make open

defs: orthogonal_complements.tex
	defs orthogonal_complements.tex

spell:
	aspell -c orthogonal_complements.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f orthogonal_complements.*
	make reset
	bbk_sheet -mode mk > Makefile