# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make numberings.pdf
	make graph.pdf

numberings.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > numberings.tex

numberings.pdf: ../../*.tex ../../trademark.pdf *.tex numberings.tex
	# pdflatex --file-line-error -interaction=nonstopmode numberings.tex
	../../latexrun numberings.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: numberings.pdf
	open numberings.pdf

o:
	make open

terms: numberings.tex
	bbk_sheet -mode ts numberings.tex

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
	rm -f numberings.*
	make reset
	bbk_sheet -mode mk > Makefile