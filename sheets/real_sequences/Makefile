# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_sequences.pdf
	make graph.pdf

real_sequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_sequences.tex

real_sequences.pdf: ../../*.tex ../../trademark.pdf *.tex real_sequences.tex
	pdflatex --file-line-error -interaction=nonstopmode real_sequences.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_sequences.pdf
	open real_sequences.pdf

o:
	make open

defs: real_sequences.tex
	defs real_sequences.tex

spell:
	aspell -c real_sequences.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm real_sequences.*
	make reset
	bbk_sheet -mode mk > Makefile