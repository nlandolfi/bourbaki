# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_continuity.pdf
	make graph.pdf

real_continuity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_continuity.tex

real_continuity.pdf: ../../*.tex ../../trademark.pdf *.tex real_continuity.tex
	pdflatex --file-line-error -interaction=nonstopmode real_continuity.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_continuity.pdf
	open real_continuity.pdf

o:
	make open

terms: real_continuity.tex
	bbk_sheet -mode ts real_continuity.tex

spell:
	aspell -c real_continuity.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_continuity.*
	make reset
	bbk_sheet -mode mk > Makefile