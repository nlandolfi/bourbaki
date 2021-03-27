# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make equations.pdf
	make graph.pdf

equations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > equations.tex

equations.pdf: ../../*.tex ../../trademark.pdf *.tex equations.tex
	pdflatex --file-line-error -interaction=nonstopmode equations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: equations.pdf
	open equations.pdf

o:
	make open

defs: equations.tex
	defs equations.tex

spell:
	aspell -c equations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f equations.*
	make reset
	bbk_sheet -mode mk > Makefile