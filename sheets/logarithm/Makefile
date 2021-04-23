# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make logarithm.pdf
	make graph.pdf

logarithm.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > logarithm.tex

logarithm.pdf: ../../*.tex ../../trademark.pdf *.tex logarithm.tex
	# pdflatex --file-line-error -interaction=nonstopmode logarithm.tex
	../../latexrun logarithm.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: logarithm.pdf
	open logarithm.pdf

o:
	make open

terms: logarithm.tex
	bbk_sheet -mode ts logarithm.tex

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
	rm -f logarithm.*
	make reset
	bbk_sheet -mode mk > Makefile