# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make equivalence_classes.pdf

equivalence_classes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > equivalence_classes.tex

equivalence_classes.pdf: ../../*.tex ../../trademark.pdf *.tex equivalence_classes.tex
	# pdflatex --file-line-error -interaction=nonstopmode equivalence_classes.tex
	../../latexrun equivalence_classes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: equivalence_classes.pdf
	open equivalence_classes.pdf

o:
	make open

terms: equivalence_classes.tex
	bbk_sheet -mode ts equivalence_classes.tex

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
	rm -f equivalence_classes.*
	make reset
	bbk_sheet -mode mk > Makefile