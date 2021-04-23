# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make monotone_classes.pdf
	make graph.pdf

monotone_classes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > monotone_classes.tex

monotone_classes.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_classes.tex
	# pdflatex --file-line-error -interaction=nonstopmode monotone_classes.tex
	../../latexrun monotone_classes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_classes.pdf
	open monotone_classes.pdf

o:
	make open

terms: monotone_classes.tex
	bbk_sheet -mode ts monotone_classes.tex

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
	rm -f monotone_classes.*
	make reset
	bbk_sheet -mode mk > Makefile