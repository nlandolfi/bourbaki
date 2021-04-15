# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make trace.pdf
	make graph.pdf

trace.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > trace.tex

trace.pdf: ../../*.tex ../../trademark.pdf *.tex trace.tex
	# pdflatex --file-line-error -interaction=nonstopmode trace.tex
	../../latexrun trace.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: trace.pdf
	open trace.pdf

o:
	make open

terms: trace.tex
	bbk_sheet -mode ts trace.tex

spell:
	aspell -c trace.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f trace.*
	make reset
	bbk_sheet -mode mk > Makefile