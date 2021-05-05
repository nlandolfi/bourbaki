# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make linear_equations.pdf
	make graph.pdf

linear_equations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_equations.tex

linear_equations.pdf: ../../*.tex ../../trademark.pdf *.tex linear_equations.tex
	# pdflatex --file-line-error -interaction=nonstopmode linear_equations.tex
	../../latexrun linear_equations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_equations.pdf
	open linear_equations.pdf

o:
	make open

terms: linear_equations.tex
	bbk_sheet -mode ts linear_equations.tex

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
	rm -f linear_equations.*
	make reset
	bbk_sheet -mode mk > Makefile