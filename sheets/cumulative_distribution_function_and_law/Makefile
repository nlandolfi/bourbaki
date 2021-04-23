# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cumulative_distribution_function_and_law.pdf
	make graph.pdf

cumulative_distribution_function_and_law.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cumulative_distribution_function_and_law.tex

cumulative_distribution_function_and_law.pdf: ../../*.tex ../../trademark.pdf *.tex cumulative_distribution_function_and_law.tex
	# pdflatex --file-line-error -interaction=nonstopmode cumulative_distribution_function_and_law.tex
	../../latexrun cumulative_distribution_function_and_law.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cumulative_distribution_function_and_law.pdf
	open cumulative_distribution_function_and_law.pdf

o:
	make open

terms: cumulative_distribution_function_and_law.tex
	bbk_sheet -mode ts cumulative_distribution_function_and_law.tex

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
	rm -f cumulative_distribution_function_and_law.*
	make reset
	bbk_sheet -mode mk > Makefile