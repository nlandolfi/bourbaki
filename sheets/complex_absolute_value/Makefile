# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_absolute_value.pdf

complex_absolute_value.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_absolute_value.tex

complex_absolute_value.pdf: ../../*.tex ../../trademark.pdf *.tex complex_absolute_value.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_absolute_value.tex
	../../latexrun complex_absolute_value.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_absolute_value.pdf
	open complex_absolute_value.pdf

o:
	make open

terms: complex_absolute_value.tex
	bbk_sheet -mode ts complex_absolute_value.tex

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
	rm -f complex_absolute_value.*
	make reset
	bbk_sheet -mode mk > Makefile