# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make tree_approximators_of_a_normal.pdf
	make graph.pdf

tree_approximators_of_a_normal.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree_approximators_of_a_normal.tex

tree_approximators_of_a_normal.pdf: ../../*.tex ../../trademark.pdf *.tex tree_approximators_of_a_normal.tex
	# pdflatex --file-line-error -interaction=nonstopmode tree_approximators_of_a_normal.tex
	../../latexrun tree_approximators_of_a_normal.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_approximators_of_a_normal.pdf
	open tree_approximators_of_a_normal.pdf

o:
	make open

terms: tree_approximators_of_a_normal.tex
	bbk_sheet -mode ts tree_approximators_of_a_normal.tex

spell:
	aspell -c tree_approximators_of_a_normal.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree_approximators_of_a_normal.*
	make reset
	bbk_sheet -mode mk > Makefile