# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make identity_functions.pdf
	make graph.pdf

identity_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > identity_functions.tex

identity_functions.pdf: ../../*.tex ../../trademark.pdf *.tex identity_functions.tex
	pdflatex --file-line-error -interaction=nonstopmode identity_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: identity_functions.pdf
	open identity_functions.pdf

o:
	make open

terms: identity_functions.tex
	bbk_sheet -mode ts identity_functions.tex

spell:
	aspell -c identity_functions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f identity_functions.*
	make reset
	bbk_sheet -mode mk > Makefile