# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make total_probability.pdf
	make graph.pdf

total_probability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > total_probability.tex

total_probability.pdf: ../../*.tex ../../trademark.pdf *.tex total_probability.tex
	pdflatex --file-line-error -interaction=nonstopmode total_probability.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: total_probability.pdf
	open total_probability.pdf

o:
	make open

defs: total_probability.tex
	defs total_probability.tex

spell:
	aspell -c total_probability.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm total_probability.*
	make reset
	bbk_sheet -mode mk > Makefile