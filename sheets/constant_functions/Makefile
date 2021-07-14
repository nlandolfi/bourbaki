# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make constant_functions.pdf

constant_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > constant_functions.tex

constant_functions.pdf: ../../*.tex ../../trademark.pdf *.tex constant_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode constant_functions.tex
	../../latexrun constant_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: constant_functions.pdf
	open constant_functions.pdf

o:
	make open

terms: constant_functions.tex
	bbk_sheet -mode ts constant_functions.tex

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
	rm -f constant_functions.*
	make reset
	bbk_sheet -mode mk > Makefile