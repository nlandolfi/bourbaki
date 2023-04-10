# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make row_reducer_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

row_reducer_matrices.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > row_reducer_matrices.tex

row_reducer_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex row_reducer_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode row_reducer_matrices.tex
	../../latexrun row_reducer_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: row_reducer_matrices.pdf
	open row_reducer_matrices.pdf

o:
	make open

terms: row_reducer_matrices.tex
	bbk sheet -mode ts row_reducer_matrices.tex

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
	rm -f row_reducer_matrices.*
	make reset
	bbk sheet -mode mk > Makefile