# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cones.pdf
	make graph.pdf

cones.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cones.tex

cones.pdf: ../../*.tex ../../trademark.pdf *.tex cones.tex
	# pdflatex --file-line-error -interaction=nonstopmode cones.tex
	../../latexrun cones.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cones.pdf
	open cones.pdf

o:
	make open

terms: cones.tex
	bbk_sheet -mode ts cones.tex

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
	rm -f cones.*
	make reset
	bbk_sheet -mode mk > Makefile