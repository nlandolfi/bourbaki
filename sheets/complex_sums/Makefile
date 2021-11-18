# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_sums.pdf

complex_sums.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_sums.tex

complex_sums.pdf: ../../*.tex ../../trademark.pdf *.tex complex_sums.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_sums.tex
	../../latexrun complex_sums.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_sums.pdf
	open complex_sums.pdf

o:
	make open

terms: complex_sums.tex
	bbk_sheet -mode ts complex_sums.tex

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
	rm -f complex_sums.*
	make reset
	bbk_sheet -mode mk > Makefile