# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make integer_arithmetic_and_order.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

integer_arithmetic_and_order.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integer_arithmetic_and_order.tex

integer_arithmetic_and_order.pdf: ../../*.tex ../../trademark.pdf *.tex integer_arithmetic_and_order.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode integer_arithmetic_and_order.tex
	../../latexrun integer_arithmetic_and_order.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_arithmetic_and_order.pdf
	open integer_arithmetic_and_order.pdf

o:
	make open

terms: integer_arithmetic_and_order.tex
	bbk_sheet -mode ts integer_arithmetic_and_order.tex

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
	rm -f integer_arithmetic_and_order.*
	make reset
	bbk_sheet -mode mk > Makefile