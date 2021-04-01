# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make identity_elements.pdf
	make graph.pdf

identity_elements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > identity_elements.tex

identity_elements.pdf: ../../*.tex ../../trademark.pdf *.tex identity_elements.tex
	pdflatex --file-line-error -interaction=nonstopmode identity_elements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: identity_elements.pdf
	open identity_elements.pdf

o:
	make open

terms: identity_elements.tex
	bbk_sheet -mode ts identity_elements.tex

spell:
	aspell -c identity_elements.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f identity_elements.*
	make reset
	bbk_sheet -mode mk > Makefile