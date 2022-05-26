# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complex_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_functions.tex

complex_functions.pdf: ../../*.tex ../../trademark.pdf *.tex complex_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_functions.tex
	../../latexrun complex_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_functions.pdf
	open complex_functions.pdf

o:
	make open

terms: complex_functions.tex
	bbk_sheet -mode ts complex_functions.tex

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
	rm -f complex_functions.*
	make reset
	bbk_sheet -mode mk > Makefile