# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make subgraphs.pdf

subgraphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > subgraphs.tex

subgraphs.pdf: ../../*.tex ../../trademark.pdf *.tex subgraphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode subgraphs.tex
	../../latexrun subgraphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: subgraphs.pdf
	open subgraphs.pdf

o:
	make open

terms: subgraphs.tex
	bbk_sheet -mode ts subgraphs.tex

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
	rm -f subgraphs.*
	make reset
	bbk_sheet -mode mk > Makefile