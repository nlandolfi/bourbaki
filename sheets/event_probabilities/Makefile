# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make event_probabilities.pdf
	make graph.pdf

event_probabilities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > event_probabilities.tex

event_probabilities.pdf: ../../*.tex ../../trademark.pdf *.tex event_probabilities.tex
	pdflatex --file-line-error -interaction=nonstopmode event_probabilities.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: event_probabilities.pdf
	open event_probabilities.pdf

o:
	make open

defs: event_probabilities.tex
	defs event_probabilities.tex

spell:
	aspell -c event_probabilities.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm event_probabilities.*
	make reset
	bbk_sheet -mode mk > Makefile