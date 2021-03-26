# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make functions.pdf
	make graph.pdf

functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > functions.tex

functions.pdf: ../../*.tex ../../trademark.pdf *.tex functions.tex
	pdflatex --file-line-error -interaction=nonstopmode functions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: functions.pdf
	open functions.pdf

o:
	make open

defs: functions.tex
	defs functions.tex

spell:
	aspell -c functions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm functions.*
	make reset
	bbk_sheet -mode mk > Makefile