# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make coins.pdf

coins.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > coins.tex

coins.pdf: ../../*.tex ../../trademark.pdf *.tex coins.tex
	# pdflatex --file-line-error -interaction=nonstopmode coins.tex
	../../latexrun coins.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: coins.pdf
	open coins.pdf

o:
	make open

terms: coins.tex
	bbk_sheet -mode ts coins.tex

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
	rm -f coins.*
	make reset
	bbk_sheet -mode mk > Makefile