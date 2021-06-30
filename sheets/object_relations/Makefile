# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make object_relations.pdf

object_relations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > object_relations.tex

object_relations.pdf: ../../*.tex ../../trademark.pdf *.tex object_relations.tex
	# pdflatex --file-line-error -interaction=nonstopmode object_relations.tex
	../../latexrun object_relations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: object_relations.pdf
	open object_relations.pdf

o:
	make open

terms: object_relations.tex
	bbk_sheet -mode ts object_relations.tex

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
	rm -f object_relations.*
	make reset
	bbk_sheet -mode mk > Makefile