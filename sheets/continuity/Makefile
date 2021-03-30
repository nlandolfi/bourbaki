# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make continuity.pdf
	make graph.pdf

continuity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > continuity.tex

continuity.pdf: ../../*.tex ../../trademark.pdf *.tex continuity.tex
	pdflatex --file-line-error -interaction=nonstopmode continuity.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: continuity.pdf
	open continuity.pdf

o:
	make open

defs: continuity.tex
	defs continuity.tex

spell:
	aspell -c continuity.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f continuity.*
	make reset
	bbk_sheet -mode mk > Makefile