# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make partitioner.pdf

partitioner.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partitioner.tex

partitioner.pdf: ../../*.tex ../../trademark.pdf *.tex partitioner.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode partitioner.tex
	../../latexrun partitioner.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partitioner.pdf
	open partitioner.pdf

o:
	make open

terms: partitioner.tex
	bbk_sheet -mode ts partitioner.tex

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
	rm -f partitioner.*
	make reset
	bbk_sheet -mode mk > Makefile