# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make differential_mutual_information_graph.pdf
	make graph.pdf

differential_mutual_information_graph.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > differential_mutual_information_graph.tex

differential_mutual_information_graph.pdf: ../../*.tex ../../trademark.pdf *.tex differential_mutual_information_graph.tex
	pdflatex --file-line-error -interaction=nonstopmode differential_mutual_information_graph.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_mutual_information_graph.pdf
	open differential_mutual_information_graph.pdf

o:
	make open

defs: differential_mutual_information_graph.tex
	defs differential_mutual_information_graph.tex

spell:
	aspell -c differential_mutual_information_graph.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f differential_mutual_information_graph.*
	make reset
	bbk_sheet -mode mk > Makefile