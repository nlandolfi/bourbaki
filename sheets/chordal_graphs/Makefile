# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make chordal_graphs.pdf

chordal_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > chordal_graphs.tex

chordal_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex chordal_graphs.tex
	# pdflatex --file-line-error -interaction=nonstopmode chordal_graphs.tex
	../../latexrun chordal_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: chordal_graphs.pdf
	open chordal_graphs.pdf

o:
	make open

terms: chordal_graphs.tex
	bbk_sheet -mode ts chordal_graphs.tex

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
	rm -f chordal_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile