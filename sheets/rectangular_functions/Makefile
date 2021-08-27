# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make rectangular_functions.pdf

rectangular_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rectangular_functions.tex

rectangular_functions.pdf: ../../*.tex ../../trademark.pdf *.tex rectangular_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode rectangular_functions.tex
	../../latexrun rectangular_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rectangular_functions.pdf
	open rectangular_functions.pdf

o:
	make open

terms: rectangular_functions.tex
	bbk_sheet -mode ts rectangular_functions.tex

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
	rm -f rectangular_functions.*
	make reset
	bbk_sheet -mode mk > Makefile