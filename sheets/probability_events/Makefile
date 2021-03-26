# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability_events.pdf
	make graph.pdf

probability_events.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_events.tex

probability_events.pdf: ../../*.tex ../../trademark.pdf *.tex probability_events.tex
	pdflatex --file-line-error -interaction=nonstopmode probability_events.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_events.pdf
	open probability_events.pdf

o:
	make open

defs: probability_events.tex
	defs probability_events.tex

spell:
	aspell -c probability_events.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm probability_events.*
	make reset
	bbk_sheet -mode mk > Makefile