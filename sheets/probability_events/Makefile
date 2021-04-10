# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability_events.pdf
	make graph.pdf

probability_events.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_events.tex

probability_events.pdf: ../../*.tex ../../trademark.pdf *.tex probability_events.tex
	# pdflatex --file-line-error -interaction=nonstopmode probability_events.tex
	../../latexrun probability_events.tex
	make terms

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

terms: probability_events.tex
	bbk_sheet -mode ts probability_events.tex

spell:
	aspell -c probability_events.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f probability_events.*
	make reset
	bbk_sheet -mode mk > Makefile