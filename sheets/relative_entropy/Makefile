# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make relative_entropy.pdf
	make graph.pdf

relative_entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > relative_entropy.tex

relative_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex relative_entropy.tex
	pdflatex --file-line-error -interaction=nonstopmode relative_entropy.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: relative_entropy.pdf
	open relative_entropy.pdf

o:
	make open

defs: relative_entropy.tex
	defs relative_entropy.tex

spell:
	aspell -c relative_entropy.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f relative_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile