# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subspaces.pdf

subspaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subspaces.tex

subspaces.pdf: ../../*.tex ../../trademark.pdf *.tex subspaces.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subspaces.tex
	../../latexrun subspaces.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subspaces.pdf
	open subspaces.pdf

o:
	make open

terms: subspaces.tex
	bbk_sheet -mode ts subspaces.tex

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
	rm -f subspaces.*
	make reset
	bbk_sheet -mode mk > Makefile