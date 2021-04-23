# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cumulative_distribution_functions.pdf
	make graph.pdf

cumulative_distribution_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cumulative_distribution_functions.tex

cumulative_distribution_functions.pdf: ../../*.tex ../../trademark.pdf *.tex cumulative_distribution_functions.tex
	# pdflatex --file-line-error -interaction=nonstopmode cumulative_distribution_functions.tex
	../../latexrun cumulative_distribution_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cumulative_distribution_functions.pdf
	open cumulative_distribution_functions.pdf

o:
	make open

terms: cumulative_distribution_functions.tex
	bbk_sheet -mode ts cumulative_distribution_functions.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f cumulative_distribution_functions.*
	make reset
	bbk_sheet -mode mk > Makefile