# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_discs.pdf

complex_discs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_discs.tex

complex_discs.pdf: ../../*.tex ../../trademark.pdf *.tex complex_discs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_discs.tex
	../../latexrun complex_discs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_discs.pdf
	open complex_discs.pdf

o:
	make open

terms: complex_discs.tex
	bbk_sheet -mode ts complex_discs.tex

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
	rm -f complex_discs.*
	make reset
	bbk_sheet -mode mk > Makefile