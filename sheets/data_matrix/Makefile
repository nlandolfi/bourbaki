# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make data_matrix.pdf

data_matrix.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > data_matrix.tex

data_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex data_matrix.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode data_matrix.tex
	../../latexrun data_matrix.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: data_matrix.pdf
	open data_matrix.pdf

o:
	make open

terms: data_matrix.tex
	bbk_sheet -mode ts data_matrix.tex

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
	rm -f data_matrix.*
	make reset
	bbk_sheet -mode mk > Makefile