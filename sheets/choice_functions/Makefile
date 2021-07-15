# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make choice_functions.pdf

choice_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > choice_functions.tex

choice_functions.pdf: ../../*.tex ../../trademark.pdf *.tex choice_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode choice_functions.tex
	../../latexrun choice_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: choice_functions.pdf
	open choice_functions.pdf

o:
	make open

terms: choice_functions.tex
	bbk_sheet -mode ts choice_functions.tex

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
	rm -f choice_functions.*
	make reset
	bbk_sheet -mode mk > Makefile