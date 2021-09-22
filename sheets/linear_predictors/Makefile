# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_predictors.pdf

linear_predictors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_predictors.tex

linear_predictors.pdf: ../../*.tex ../../trademark.pdf *.tex linear_predictors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_predictors.tex
	../../latexrun linear_predictors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_predictors.pdf
	open linear_predictors.pdf

o:
	make open

terms: linear_predictors.tex
	bbk_sheet -mode ts linear_predictors.tex

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
	rm -f linear_predictors.*
	make reset
	bbk_sheet -mode mk > Makefile