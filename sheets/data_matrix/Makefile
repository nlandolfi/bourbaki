# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make data_matrix.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

data_matrix.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > data_matrix.tex

data_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex data_matrix.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode data_matrix.tex
	../../latexrun data_matrix.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: data_matrix.pdf
	open data_matrix.pdf

o:
	make open

terms: data_matrix.tex
	bbk sheet -mode ts data_matrix.tex

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
	bbk sheet -mode mk > Makefile