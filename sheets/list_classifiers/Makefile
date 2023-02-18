# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make list_classifiers.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

list_classifiers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > list_classifiers.tex

list_classifiers.pdf: ../../*.tex ../../trademark.pdf *.tex list_classifiers.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode list_classifiers.tex
	../../latexrun list_classifiers.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: list_classifiers.pdf
	open list_classifiers.pdf

o:
	make open

terms: list_classifiers.tex
	bbk_sheet -mode ts list_classifiers.tex

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
	rm -f list_classifiers.*
	make reset
	bbk_sheet -mode mk > Makefile