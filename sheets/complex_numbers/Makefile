# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make complex_numbers.pdf
	make graph.pdf

complex_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_numbers.tex

complex_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex complex_numbers.tex
	pdflatex --file-line-error -interaction=nonstopmode complex_numbers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_numbers.pdf
	open complex_numbers.pdf

o:
	make open

terms: complex_numbers.tex
	bbk_sheet -mode ts complex_numbers.tex

spell:
	aspell -c complex_numbers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f complex_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile