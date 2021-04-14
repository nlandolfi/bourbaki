# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make tree_normal_approximators.pdf
	make graph.pdf

tree_normal_approximators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree_normal_approximators.tex

tree_normal_approximators.pdf: ../../*.tex ../../trademark.pdf *.tex tree_normal_approximators.tex
	# pdflatex --file-line-error -interaction=nonstopmode tree_normal_approximators.tex
	../../latexrun tree_normal_approximators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_normal_approximators.pdf
	open tree_normal_approximators.pdf

o:
	make open

terms: tree_normal_approximators.tex
	bbk_sheet -mode ts tree_normal_approximators.tex

spell:
	aspell -c tree_normal_approximators.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree_normal_approximators.*
	make reset
	bbk_sheet -mode mk > Makefile