# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_analytic_functions.pdf

complex_analytic_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_analytic_functions.tex

complex_analytic_functions.pdf: ../../*.tex ../../trademark.pdf *.tex complex_analytic_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_analytic_functions.tex
	../../latexrun complex_analytic_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_analytic_functions.pdf
	open complex_analytic_functions.pdf

o:
	make open

terms: complex_analytic_functions.tex
	bbk_sheet -mode ts complex_analytic_functions.tex

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
	rm -f complex_analytic_functions.*
	make reset
	bbk_sheet -mode mk > Makefile