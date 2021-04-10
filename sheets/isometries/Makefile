# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make isometries.pdf
	make graph.pdf

isometries.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > isometries.tex

isometries.pdf: ../../*.tex ../../trademark.pdf *.tex isometries.tex
	# pdflatex --file-line-error -interaction=nonstopmode isometries.tex
	../../latexrun isometries.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: isometries.pdf
	open isometries.pdf

o:
	make open

terms: isometries.tex
	bbk_sheet -mode ts isometries.tex

spell:
	aspell -c isometries.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f isometries.*
	make reset
	bbk_sheet -mode mk > Makefile