# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make norm_examples.pdf
	make graph.pdf

norm_examples.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > norm_examples.tex

norm_examples.pdf: ../../*.tex ../../trademark.pdf *.tex norm_examples.tex
	pdflatex --file-line-error -interaction=nonstopmode norm_examples.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norm_examples.pdf
	open norm_examples.pdf

o:
	make open

defs: norm_examples.tex
	defs norm_examples.tex

spell:
	aspell -c norm_examples.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f norm_examples.*
	make reset
	bbk_sheet -mode mk > Makefile