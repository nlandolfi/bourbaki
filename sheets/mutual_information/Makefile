# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make mutual_information.pdf
	make graph.pdf

mutual_information.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > mutual_information.tex

mutual_information.pdf: ../../*.tex ../../trademark.pdf *.tex mutual_information.tex
	pdflatex --file-line-error -interaction=nonstopmode mutual_information.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: mutual_information.pdf
	open mutual_information.pdf

o:
	make open

defs: mutual_information.tex
	defs mutual_information.tex

spell:
	aspell -c mutual_information.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f mutual_information.*
	make reset
	bbk_sheet -mode mk > Makefile