# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make classifiers.pdf
	make graph.pdf

classifiers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > classifiers.tex

classifiers.pdf: ../../*.tex ../../trademark.pdf *.tex classifiers.tex
	# pdflatex --file-line-error -interaction=nonstopmode classifiers.tex
	../../latexrun classifiers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: classifiers.pdf
	open classifiers.pdf

o:
	make open

terms: classifiers.tex
	bbk_sheet -mode ts classifiers.tex

spell:
	aspell -c classifiers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f classifiers.*
	make reset
	bbk_sheet -mode mk > Makefile