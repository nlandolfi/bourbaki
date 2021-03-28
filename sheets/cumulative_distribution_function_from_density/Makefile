# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cumulative_distribution_function_from_density.pdf
	make graph.pdf

cumulative_distribution_function_from_density.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cumulative_distribution_function_from_density.tex

cumulative_distribution_function_from_density.pdf: ../../*.tex ../../trademark.pdf *.tex cumulative_distribution_function_from_density.tex
	pdflatex --file-line-error -interaction=nonstopmode cumulative_distribution_function_from_density.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cumulative_distribution_function_from_density.pdf
	open cumulative_distribution_function_from_density.pdf

o:
	make open

defs: cumulative_distribution_function_from_density.tex
	defs cumulative_distribution_function_from_density.tex

spell:
	aspell -c cumulative_distribution_function_from_density.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f cumulative_distribution_function_from_density.*
	make reset
	bbk_sheet -mode mk > Makefile