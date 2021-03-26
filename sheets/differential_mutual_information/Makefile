# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make differential_mutual_information.pdf
	make graph.pdf

differential_mutual_information.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > differential_mutual_information.tex

differential_mutual_information.pdf: ../../*.tex ../../trademark.pdf *.tex differential_mutual_information.tex
	pdflatex --file-line-error -interaction=nonstopmode differential_mutual_information.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differential_mutual_information.pdf
	open differential_mutual_information.pdf

o:
	make open

defs: differential_mutual_information.tex
	defs differential_mutual_information.tex

spell:
	aspell -c differential_mutual_information.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm differential_mutual_information.*
	make reset
	bbk_sheet -mode mk > Makefile