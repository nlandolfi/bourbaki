# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make vectors.pdf
	make graph.pdf

vectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > vectors.tex

vectors.pdf: ../../*.tex ../../trademark.pdf *.tex vectors.tex
	pdflatex --file-line-error -interaction=nonstopmode vectors.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vectors.pdf
	open vectors.pdf

o:
	make open

defs: vectors.tex
	defs vectors.tex

spell:
	aspell -c vectors.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f vectors.*
	make reset
	bbk_sheet -mode mk > Makefile