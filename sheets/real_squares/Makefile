# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_squares.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_squares.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_squares.tex

real_squares.pdf: ../../*.tex ../../trademark.pdf *.tex real_squares.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_squares.tex
	../../latexrun real_squares.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_squares.pdf
	open real_squares.pdf

o:
	make open

terms: real_squares.tex
	bbk_sheet -mode ts real_squares.tex

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
	rm -f real_squares.*
	make reset
	bbk_sheet -mode mk > Makefile