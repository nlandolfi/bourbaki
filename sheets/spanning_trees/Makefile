# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make spanning_trees.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

spanning_trees.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > spanning_trees.tex

spanning_trees.pdf: ../../*.tex ../../trademark.pdf *.tex spanning_trees.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode spanning_trees.tex
	../../latexrun spanning_trees.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: spanning_trees.pdf
	open spanning_trees.pdf

o:
	make open

terms: spanning_trees.tex
	bbk sheet -mode ts spanning_trees.tex

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
	rm -f spanning_trees.*
	make reset
	bbk sheet -mode mk > Makefile