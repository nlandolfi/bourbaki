# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make bounded_functions.pdf

bounded_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > bounded_functions.tex

bounded_functions.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_functions.tex
	# pdflatex --file-line-error -interaction=nonstopmode bounded_functions.tex
	../../latexrun bounded_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: bounded_functions.pdf
	open bounded_functions.pdf

o:
	make open

terms: bounded_functions.tex
	bbk_sheet -mode ts bounded_functions.tex

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
	rm -f bounded_functions.*
	make reset
	bbk_sheet -mode mk > Makefile