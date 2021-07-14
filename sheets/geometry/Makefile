# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make geometry.pdf

geometry.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > geometry.tex

geometry.pdf: ../../*.tex ../../trademark.pdf *.tex geometry.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode geometry.tex
	../../latexrun geometry.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: geometry.pdf
	open geometry.pdf

o:
	make open

terms: geometry.tex
	bbk_sheet -mode ts geometry.tex

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
	rm -f geometry.*
	make reset
	bbk_sheet -mode mk > Makefile