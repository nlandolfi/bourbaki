# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make perfect_elimination_orderings.pdf

perfect_elimination_orderings.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > perfect_elimination_orderings.tex

perfect_elimination_orderings.pdf: ../../*.tex ../../trademark.pdf *.tex perfect_elimination_orderings.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode perfect_elimination_orderings.tex
	../../latexrun perfect_elimination_orderings.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: perfect_elimination_orderings.pdf
	open perfect_elimination_orderings.pdf

o:
	make open

terms: perfect_elimination_orderings.tex
	bbk_sheet -mode ts perfect_elimination_orderings.tex

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
	rm -f perfect_elimination_orderings.*
	make reset
	bbk_sheet -mode mk > Makefile