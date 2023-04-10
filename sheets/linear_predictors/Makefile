# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_predictors.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

linear_predictors.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > linear_predictors.tex

linear_predictors.pdf: ../../*.tex ../../trademark.pdf *.tex linear_predictors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_predictors.tex
	../../latexrun linear_predictors.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_predictors.pdf
	open linear_predictors.pdf

o:
	make open

terms: linear_predictors.tex
	bbk sheet -mode ts linear_predictors.tex

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
	bbk sheet -mode mk > Makefile