# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make event_independence.pdf
	make graph.pdf

event_independence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > event_independence.tex

event_independence.pdf: ../../*.tex ../../trademark.pdf *.tex event_independence.tex
	pdflatex --file-line-error -interaction=nonstopmode event_independence.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: event_independence.pdf
	open event_independence.pdf

o:
	make open

defs: event_independence.tex
	defs event_independence.tex

spell:
	aspell -c event_independence.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f event_independence.*
	make reset
	bbk_sheet -mode mk > Makefile