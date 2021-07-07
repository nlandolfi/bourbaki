# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distribution_expectation.pdf

distribution_expectation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_expectation.tex

distribution_expectation.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_expectation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode distribution_expectation.tex
	../../latexrun distribution_expectation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_expectation.pdf
	open distribution_expectation.pdf

o:
	make open

terms: distribution_expectation.tex
	bbk_sheet -mode ts distribution_expectation.tex

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
	rm -f distribution_expectation.*
	make reset
	bbk_sheet -mode mk > Makefile