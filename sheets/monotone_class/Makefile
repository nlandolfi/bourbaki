# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make monotone_class.pdf
	make graph.pdf

monotone_class.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monotone_class.tex

monotone_class.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_class.tex
	pdflatex --file-line-error -interaction=nonstopmode monotone_class.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_class.pdf
	open monotone_class.pdf

o:
	make open

defs: monotone_class.tex
	defs monotone_class.tex

spell:
	aspell -c monotone_class.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f monotone_class.*
	make reset
	bbk_sheet -mode mk > Makefile