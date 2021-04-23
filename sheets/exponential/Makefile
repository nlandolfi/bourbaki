# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make exponential.pdf
	make graph.pdf

exponential.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > exponential.tex

exponential.pdf: ../../*.tex ../../trademark.pdf *.tex exponential.tex
	# pdflatex --file-line-error -interaction=nonstopmode exponential.tex
	../../latexrun exponential.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: exponential.pdf
	open exponential.pdf

o:
	make open

terms: exponential.tex
	bbk_sheet -mode ts exponential.tex

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
	rm -f exponential.*
	make reset
	bbk_sheet -mode mk > Makefile