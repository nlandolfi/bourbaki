# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make exponential_derivatives.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

exponential_derivatives.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > exponential_derivatives.tex

exponential_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex exponential_derivatives.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode exponential_derivatives.tex
	../../latexrun exponential_derivatives.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: exponential_derivatives.pdf
	open exponential_derivatives.pdf

o:
	make open

terms: exponential_derivatives.tex
	bbk sheet -mode ts exponential_derivatives.tex

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
	rm -f exponential_derivatives.*
	make reset
	bbk sheet -mode mk > Makefile