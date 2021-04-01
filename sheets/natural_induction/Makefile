# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_induction.pdf
	make graph.pdf

natural_induction.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_induction.tex

natural_induction.pdf: ../../*.tex ../../trademark.pdf *.tex natural_induction.tex
	pdflatex --file-line-error -interaction=nonstopmode natural_induction.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_induction.pdf
	open natural_induction.pdf

o:
	make open

terms: natural_induction.tex
	bbk_sheet -mode ts natural_induction.tex

spell:
	aspell -c natural_induction.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f natural_induction.*
	make reset
	bbk_sheet -mode mk > Makefile