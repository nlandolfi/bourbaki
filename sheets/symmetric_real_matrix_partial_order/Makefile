# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make symmetric_real_matrix_partial_order.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

symmetric_real_matrix_partial_order.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > symmetric_real_matrix_partial_order.tex

symmetric_real_matrix_partial_order.pdf: ../../*.tex ../../trademark.pdf *.tex symmetric_real_matrix_partial_order.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode symmetric_real_matrix_partial_order.tex
	../../latexrun symmetric_real_matrix_partial_order.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: symmetric_real_matrix_partial_order.pdf
	open symmetric_real_matrix_partial_order.pdf

o:
	make open

terms: symmetric_real_matrix_partial_order.tex
	bbk_sheet -mode ts symmetric_real_matrix_partial_order.tex

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
	rm -f symmetric_real_matrix_partial_order.*
	make reset
	bbk_sheet -mode mk > Makefile