# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make independent_identical_sums.pdf
	make graph.pdf

independent_identical_sums.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > independent_identical_sums.tex

independent_identical_sums.pdf: ../../*.tex ../../trademark.pdf *.tex independent_identical_sums.tex
	pdflatex --file-line-error -interaction=nonstopmode independent_identical_sums.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: independent_identical_sums.pdf
	open independent_identical_sums.pdf

o:
	make open

defs: independent_identical_sums.tex
	defs independent_identical_sums.tex

spell:
	aspell -c independent_identical_sums.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f independent_identical_sums.*
	make reset
	bbk_sheet -mode mk > Makefile